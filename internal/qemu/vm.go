package qemu

import (
	"bytes"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"github.com/CharVstack/CharV-backend/domain/models"
	"github.com/mattn/go-shellwords"
)

func install(opts InstallOpts, filePath string) (models.Vm, error) {
	tmpl, err := template.New("install").Parse(`qemu-system-x86_64 -accel kvm -daemonize -display none -name guest={{.Name}} -smp {{.VCpu}} -m {{.Memory}} -cdrom /var/lib/charVstack/iso/{{.Image}} -boot order=d -drive file=/var/lib/charVstack/images/{{.Disk}}.qcow2,format=qcow2 -drive file=/var/lib/charVstack/bios/bios.bin,format=raw,if=pflash,readonly=on`)
	if err != nil {
		return models.Vm{}, err
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, opts)
	if err != nil {
		return models.Vm{}, err
	}
	cmd := buf.String()

	var resJSON models.Vm
	resJSON, err = createInfoJSON(opts, filePath)
	if err != nil {
		return models.Vm{}, err
	}

	return resJSON, run(cmd)
}

// CreateVm diskとVmをcharV-libの関数から作成する
func CreateVm(vmInfo models.PostApiV1VmsJSONRequestBody) (models.Vm, error) {
	toGetInfo := InstallOpts{
		Name:   vmInfo.Name,
		Memory: vmInfo.Memory,
		VCpu:   vmInfo.Vcpu,
		Image:  "ubuntu-20.04.5-live-server-amd64.iso",
		Disk:   vmInfo.Name + "Disk",
	}

	name, err := createDisk(toGetInfo.Disk)
	if err != nil {
		return models.Vm{}, err
	}

	var createVm models.Vm
	createVm, err = install(toGetInfo, name)
	if err != nil {
		return models.Vm{}, err
	}

	return createVm, err
}

func run(cmd string) error {
	c, err := shellwords.Parse(cmd)
	if err != nil {
		return err
	}
	switch len(c) {
	case 0:
		return nil
	case 1:
		err = exec.Command(c[0]).Run()
	default:
		err = exec.Command(c[0], c[1:]...).Run()
	}
	if err != nil {
		return err
	}
	return nil
}

func parse(path string) (models.Vm, error) {
	var machine models.Vm
	abspath, err := filepath.Abs(path)
	raw, err := os.ReadFile(abspath)
	if err != nil {
		return models.Vm{}, err
	}
	if err := json.Unmarshal(raw, &machine); err != nil {
		return models.Vm{}, err
	}
	return machine, nil
}
