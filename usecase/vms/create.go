package vms

import (
	"github.com/CharVstack/CharV-backend/openapi/v1"
	"github.com/CharVstack/CharV-lib/qemu"
)

// CreateVm diskとVmをcharV-libの関数から作成する
func CreateVm(vmInfo openapi.PostApiV1VmsJSONRequestBody) (error, error) {
	getVmInfo := qemu.InstallOpts{
		Name:   vmInfo.Name,
		Memory: vmInfo.Memory,
		VCpu:   vmInfo.Vcpu,
		Image:  "ubuntu-20.04.5-live-server-amd64.iso",
		Disk:   vmInfo.Name + "disk",
	}

	createDisk := qemu.CreateDisk(getVmInfo.Disk)
	createVm := qemu.Install(getVmInfo)

	return createDisk, createVm
}
