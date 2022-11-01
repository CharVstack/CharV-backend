package handler

import (
	"net/http"
	"os"

	"github.com/CharVstack/CharV-lib/pkg/qemu"

	backendHost "github.com/CharVstack/CharV-backend/usecase/host"
	"github.com/CharVstack/CharV-backend/usecase/vms"
	"github.com/CharVstack/CharV-lib/domain/models"
	libHost "github.com/CharVstack/CharV-lib/pkg/host"
	"github.com/gin-gonic/gin"
)

// V1Handler 引数を返さないので空
type V1Handler struct{}

func (v V1Handler) GetApiV1Host(c *gin.Context) {
	storageDirEnv := os.Getenv("STORAGE_DIR")
	storageDir := models.GetInfoOptions{
		StorageDir: storageDirEnv,
	}
	getInfo, err := libHost.GetInfo(storageDir)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	hostInfo := backendHost.GetHostInfo(getInfo)
	c.JSON(http.StatusOK, hostInfo)
}

func (v V1Handler) GetApiV1Vms(c *gin.Context) {
	vmsInfo, err := vms.Info()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"vms": vmsInfo,
	})
}

// PostApiV1Vms Vm作成時にフロントから情報を受取りステータスを返す
func (v V1Handler) PostApiV1Vms(c *gin.Context) {
	var requestBody models.PostApiV1VmsJSONRequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := qemu.ExistsSufficientMemory(uint64(requestBody.Memory))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var vm models.Vm
	vm, err = vms.CreateVm(requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, vm)
}

func (v V1Handler) GetApiV1VmsVmId(c *gin.Context, vmId string) {
	//TODO implement me
	panic("implement me")
}

func (v V1Handler) PatchApiV1VmsVmId(c *gin.Context, vmId string) {
	//TODO implement me
	panic("implement me")
}
