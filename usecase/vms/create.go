package vms

import (
	"github.com/CharVstack/CharV-lib/domain/models"
	"github.com/CharVstack/CharV-lib/pkg/qemu"
)

// CreateVm diskとVmをcharV-libの関数から作成する
func CreateVm(vmInfo models.PostApiV1VmsJSONRequestBody) (models.Vm, error) {
	toGetInfo := models.InstallOpts{
		Name:   vmInfo.Name,
		Memory: vmInfo.Memory,
		VCpu:   vmInfo.Vcpu,
		Image:  "ubuntu-20.04.5-live-server-amd64.iso",
		Disk:   vmInfo.Name + "Disk",
	}

	name, err := qemu.CreateDisk(toGetInfo.Disk)
	if err != nil {
		return models.Vm{}, err
	}

	var createVm models.Vm{}
	createVm, err = qemu.Install(toGetInfo, name)
	if err != nil {
		return models.Vm{}, err
	}

	return createVm, err
}
