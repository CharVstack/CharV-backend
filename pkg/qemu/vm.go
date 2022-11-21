package qemu

import (
	"bytes"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"github.com/CharVstack/CharV-backend/domain/models"
	"github.com/google/uuid"
	"github.com/mattn/go-shellwords"
)

func install(opts InstallOpts, filePath string) (models.Vm, error) {
	tmpl, err := template.New("install").Parse(`qemu-system-x86_64 -accel kvm -daemonize -display none -name guest={{.Name}} -smp {{.VCpu}} -m {{.Memory}} -cdrom /var/lib/charVstack/iso/{{.Image}} -boot order=d -drive file=/var/lib/charVstack/images/{{.Disk}}.qcow2,format=qcow2 -drive file=/var/lib/charVstack/bios/bios.bin,format=raw,if=pflash,readonly=on -monitor unix:/{{.SocketPath}}/{{.Id}}.sock,server,nowait`)
	if err != nil {
		return models.Vm{}, err
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, opts)
	if err != nil {
		return models.Vm{}, err
	}
	cmd := buf.String()

	vm, err := getVmInfo(opts, filePath)
	if err != nil {
		return models.Vm{}, err
	}

	return vm, run(cmd)
}

// CreateVm diskとVmをcharV-libの関数から作成する
func CreateVm(vmInfo models.PostApiV1VmsJSONRequestBody, socksPath string) (models.Vm, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return models.Vm{}, err
	}

	toGetInfo := InstallOpts{
		Name:       vmInfo.Name,
		Memory:     vmInfo.Memory,
		VCpu:       vmInfo.Cpu,
		Image:      "ubuntu-20.04.5-live-server-amd64.iso",
		Disk:       vmInfo.Name + "Disk", // ToDo: 現在ストレージの複数作成機能がフロントエンドにないので１つを前提にしている
		Id:         id,
		SocketPath: socksPath,
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

func getAllVms(directoryPath string) ([]models.Vm, error) {
	var vmList []models.Vm
	dir, err := os.ReadDir(directoryPath)
	if err != nil {
		return []models.Vm{}, err
	}

	for _, file := range dir {
		raw, err := os.ReadFile(directoryPath + file.Name())
		if err != nil {
			return []models.Vm{}, err
		}

		var vm models.Vm
		err = json.Unmarshal(raw, &vm)
		if err != nil {
			return []models.Vm{}, err
		}

		vmList = append(vmList, vm)

	}
	return vmList, err
}

func getVmInfo(opts InstallOpts, filePath string) (vmInfo models.Vm, err error) {
	if err != nil {
		return models.Vm{}, err
	}

	var diskType models.DiskType
	diskType, err = CheckFileType(filePath)
	if err != nil {
		return models.Vm{}, err
	}

	vmInfo = models.Vm{
		Devices: models.Devices{
			Disk: []models.Disk{
				{
					Type:   diskType,
					Device: models.DiskDeviceDisk,
					Name:   opts.Disk,
					Pool:   "default", // ToDo: ストレージの選択機能がないので "default" で固定
				},
			},
		},
		Memory: opts.Memory,
		Metadata: models.Metadata{
			ApiVersion: "v1",
			Id:         opts.Id,
		},
		Name: opts.Name,
		Cpu:  opts.VCpu,
	}

	rawVmInfo, err := json.Marshal(vmInfo)
	if err != nil {
		return models.Vm{}, err
	}

	createJSONPath := "/var/lib/charVstack/machines/"

	fileName := createJSONPath + vmInfo.Name + "-" + vmInfo.Metadata.Id.String() + ".json"

	createFile, err := os.Create(fileName)
	if err != nil {
		return models.Vm{}, err
	}
	defer func() {
		err = createFile.Close()
	}()

	_, err = createFile.Write(rawVmInfo)
	if err != nil {
		return models.Vm{}, err
	}

	return vmInfo, nil
}

func GetAllVmInfo() ([]models.Vm, error) {
	path := "/var/lib/charVstack/machines/"
	vms, err := getAllVms(path)
	if err != nil {
		return []models.Vm{}, err
	}
	return vms, nil
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
