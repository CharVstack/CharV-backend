package vms

import (
	"fmt"
	"github.com/CharVstack/CharV-backend/openapi/v1"
	"github.com/CharVstack/CharV-lib/qemu"
)

func GetVmsInfo() ([]string, openapi.Vm) {
	list, err := qemu.GetRunningList()
	if err != nil {
		fmt.Println("testQEMU ERROR")
		return nil, openapi.Vm{}
	}

	vmInfo := openapi.Vm{
		Devices: struct {
			Disk []struct {
				Path string `json:"path"`
				Type string `json:"type"`
			} `json:"disk"`
		}{},
		Memory: 0,
		Metadata: struct {
			ApiVersion string `json:"api_version"`
			Id         string `json:"id"`
		}{},
		Name: "",
		Vcpu: 0,
	}

	return list, vmInfo
}
