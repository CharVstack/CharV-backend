package handler

import (
	"net/http"
	"os"

	"github.com/CharVstack/CharV-backend/openapi/v1"
	backendHost "github.com/CharVstack/CharV-backend/usecase/host"
	"github.com/CharVstack/CharV-backend/usecase/vms"
	"github.com/CharVstack/CharV-lib/pkg/host"
	"github.com/gin-gonic/gin"
)

// V1Handler 引数を返さないので空
type V1Handler struct{}

func (v V1Handler) GetApiV1Host(c *gin.Context) {
	storageDirEnv := os.Getenv("STORAGE_DIR")
	storageDir := host.GetInfoOptions{
		StorageDir: storageDirEnv,
	}
	getInfo := host.GetInfo(storageDir)
	hostInfo := backendHost.GetHostInfo(getInfo)
	c.JSON(http.StatusOK, gin.H{
		"cpu":           hostInfo.Cpu,
		"mem":           hostInfo.Mem,
		"storage_pools": hostInfo.StoragePools,
	})
}

func (v V1Handler) GetApiV1Vms(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

// PostApiV1Vms Vm作成時にフロントから情報を受取りステータスを返す
func (v V1Handler) PostApiV1Vms(c *gin.Context) {
	var getJsonData openapi.PostApiV1VmsJSONRequestBody
	if err := c.ShouldBindJSON(&getJsonData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createDiskErr, createVmErr := vms.CreateVm(getJsonData)
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
