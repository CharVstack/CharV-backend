package handler

import (
	"net/http"
	"os"

	"github.com/CharVstack/CharV-lib/domain"

	"github.com/CharVstack/CharV-backend/domain/models"
	backendHost "github.com/CharVstack/CharV-backend/usecase/host"
	"github.com/CharVstack/CharV-backend/usecase/vms"
	"github.com/CharVstack/CharV-lib/pkg/host"
	"github.com/gin-gonic/gin"
)

// V1Handler 引数を返さないので空
type V1Handler struct{}

func (v V1Handler) GetApiV1Host(c *gin.Context) {
	storageDirEnv := os.Getenv("STORAGE_DIR")
	storageDir := domain.GetInfoOptions{
		StorageDir: storageDirEnv,
	}
	hostInfo, err := host.GetInfo(storageDir)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": "not get hostInfo",
		})
	}
	res := backendHost.GetHostInfo(hostInfo)
	c.JSON(http.StatusOK, gin.H{
		"cpu":           res.Cpu,
		"mem":           res.Mem,
		"storage_pools": res.StoragePools,
	})
}

func (v V1Handler) GetApiV1Vms(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

// PostApiV1Vms Vm作成時にフロントから情報を受取りステータスを返す
func (v V1Handler) PostApiV1Vms(c *gin.Context) {
	var req models.PostApiV1VmsJSONRequestBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vmInfo, createDiskErr, createVmErr := vms.CreateVm(req)
	if createDiskErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error: Disk not created": createDiskErr.Error()})
		return
	}
	if createVmErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error: Vm not created": createVmErr.Error()})
		return
	}

	responseData := models.Vm{
		Devices: struct {
			Disk []struct {
				Path string `json:"path"`
				Type string `json:"type"`
			} `json:"disk"`
		}{
			Disk: []struct {
				Path string `json:"path"`
				Type string `json:"type"`
			}{
				{
					Path: vmInfo.Devices.Disk[0].Path,
					Type: vmInfo.Devices.Disk[0].Type,
				},
			},
		},
		Memory: vmInfo.Memory,
		Metadata: struct {
			ApiVersion string `json:"api_version"`
			Id         string `json:"id"`
		}{
			ApiVersion: vmInfo.Metadata.ApiVersion,
			Id:         vmInfo.Metadata.Id,
		},
		Name: vmInfo.Name,
		Vcpu: vmInfo.VCpu,
	}

	c.JSON(http.StatusOK, gin.H{
		"name":     responseData.Name,
		"metadata": responseData.Metadata,
		"memory":   responseData.Memory,
		"vcpu":     responseData.Vcpu,
		"devices":  responseData.Devices,
	})
}

func (v V1Handler) GetApiV1VmsVmId(c *gin.Context, vmId string) {
	//TODO implement me
	panic("implement me")
}

func (v V1Handler) PatchApiV1VmsVmId(c *gin.Context, vmId string) {
	//TODO implement me
	panic("implement me")
}
