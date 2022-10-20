package vms

import (
	"github.com/CharVstack/CharV-backend/domain/models"
	"github.com/CharVstack/CharV-lib/domain"
	"github.com/CharVstack/CharV-lib/pkg/qemu"
)

// CreateVm diskとVmをcharV-libの関数から作成する
func CreateVm(vmInfo models.PostApiV1VmsJSONRequestBody) (domain.Vm, error, error) {
	getVmInfo := domain.InstallOpts{
		Name:   vmInfo.Name,
		Memory: vmInfo.Memory,
		VCpu:   vmInfo.Vcpu,
		Image:  "ubuntu-20.04.5-live-server-amd64.iso",
		Disk:   vmInfo.Name + "disk",
	}

	createDisk := qemu.CreateDisk(getVmInfo.Disk)
	getJSONData, createVm := qemu.Install(getVmInfo)

	return getJSONData, createDisk, createVm
}
