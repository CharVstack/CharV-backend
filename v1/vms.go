package openapi

import (
	"fmt"
	"github.com/CharVstack/CharV-lib/qemu"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetVmsInfo() ([]string, Vm) {
	list, err := qemu.GetRunningList()
	if err != nil {
		fmt.Println("testQEMU ERROR")
		return nil, Vm{}
	}

	vmInfo := Vm{
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

func (v V1Handler) GetApiV1Vms(c *gin.Context) {
	getVmTestData, getVmsInfo := GetVmsInfo()
	fmt.Println("getVmsInfo: ", getVmsInfo)
	//TODO implement me
	c.JSON(http.StatusOK, gin.H{
		"message": getVmTestData,
		//"vms":     getVmsInfo,
	})
}

func (v V1Handler) PostApiV1Vms(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (v V1Handler) GetApiV1VmsVmId(c *gin.Context, vmId string) {
	//TODO implement me
	panic("implement me")
}

func (v V1Handler) PatchApiV1VmsVmId(c *gin.Context, vmId string) {
	//TODO implement me
	panic("implement me")
}
