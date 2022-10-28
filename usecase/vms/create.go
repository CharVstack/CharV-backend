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
		Disk:   vmInfo.Name + "disk",
	}

	if err := qemu.CreateDisk(toGetInfo.Disk); err != nil {
		return models.Vm{}, err
	}

	createVm, err := qemu.Install(toGetInfo)
	if err != nil {
		return models.Vm{}, err
	}

	return createVm, err
}
