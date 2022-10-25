package vms

import (
	backendModels "github.com/CharVstack/CharV-backend/domain/models"
	libModels "github.com/CharVstack/CharV-lib/domain/models"
	"github.com/CharVstack/CharV-lib/pkg/qemu"
)

// CreateVm diskとVmをcharV-libの関数から作成する
func CreateVm(vmInfo backendModels.PostApiV1VmsJSONRequestBody) (libModels.Vm, error, error) {
	getVmInfo := libModels.InstallOpts{
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
