package qemu

import (
	"bytes"
	"errors"
	"os"
	"text/template"

	"github.com/CharVstack/CharV-backend/domain/models"
)

func createDisk(name string) (string, error) {
	name = "/var/lib/charVstack/images/" + name + "." + "qcow2"
	tmpl, err := template.New("create").Parse(`qemu-img create -f qcow2 {{.}} 16G`)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, name)
	if err != nil {
		return "", err
	}
	cmd := buf.String()

	return name, run(cmd)
}

func CheckFileType(filePath string) (models.DiskType, error) {
	buf, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	if !bytes.Equal(buf[:4], []byte("QFI\xfb")) {
		return "", errors.New("Not QEMU QCOW Image (v3) ")
	}

	return models.DiskTypeQcow2, nil
}
