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
		"vms":     getVmsInfo,
	})
}

// CreateVm diskとVmをcharV-libの関数から作成する
func CreateVm(vmInfo PostApiV1VmsJSONRequestBody) (error, error) {
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

// PostApiV1Vms Vm作成時にフロントから情報を受取りステータスを返す
func (v V1Handler) PostApiV1Vms(c *gin.Context) {
	var getJsonData PostApiV1VmsJSONRequestBody
	if err := c.ShouldBindJSON(&getJsonData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createDiskErr, createVmErr := CreateVm(getJsonData)
	if createDiskErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error: Disk not created": createDiskErr.Error()})
		return
	}
	if createVmErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error: Vm not created": createVmErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (v V1Handler) GetApiV1VmsVmId(c *gin.Context, vmId string) {
	//TODO implement me
	panic("implement me")
}

func (v V1Handler) PatchApiV1VmsVmId(c *gin.Context, vmId string) {
	//TODO implement me
	panic("implement me")
}
