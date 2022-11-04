package handler

import (
	"net/http"
	"os"

	"github.com/CharVstack/CharV-backend/internal/util"

	"github.com/CharVstack/CharV-backend/domain/models"
	"github.com/CharVstack/CharV-backend/internal/qemu"
	"github.com/CharVstack/CharV-backend/pkg/host"
	"github.com/gin-gonic/gin"
)

// V1Handler 引数を返さないので空
type V1Handler struct{}

func (v V1Handler) GetApiV1Host(c *gin.Context) {
	storageDirEnv := os.Getenv("STORAGE_DIR")
	storageDir := host.GetInfoOptions{
		StorageDir: storageDirEnv,
	}
	hostInfo, err := host.GetInfo(storageDir)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, models.GetHost200Response{
		Host: hostInfo,
	})
}

func (v V1Handler) GetApiV1Vms(c *gin.Context) {
	allVmsInfo, err := qemu.GetAllVmInfo()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.GetAllVMsList200Response{
		Vms: allVmsInfo,
	})
}

// PostApiV1Vms Vm作成時にフロントから情報を受取りステータスを返す
func (v V1Handler) PostApiV1Vms(c *gin.Context) {
	var requestBody models.PostApiV1VmsJSONRequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := util.ExistsSufficientMemory(uint64(requestBody.Memory))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var vm models.Vm
	vm, err = qemu.CreateVm(requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.PostCreateNewVM200Response{
		Vm: vm,
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
