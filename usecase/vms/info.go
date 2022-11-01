package vms

import (
	"github.com/CharVstack/CharV-lib/domain/models"
	"github.com/CharVstack/CharV-lib/pkg/qemu"
)

func Info() ([]models.Vm, error) {
	path := "/var/lib/charVstack/machines/"
	vms, err := qemu.ConvertToStruct(path)
	if err != nil {
		return []models.Vm{}, err
	}
	return vms, nil
}
