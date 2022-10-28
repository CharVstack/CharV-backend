package handler

import (
	"net/http"
	"os"

	backendHost "github.com/CharVstack/CharV-backend/usecase/host"
	"github.com/CharVstack/CharV-backend/usecase/vms"
	"github.com/CharVstack/CharV-lib/domain/models"
	"github.com/CharVstack/CharV-lib/pkg/host"
	"github.com/gin-gonic/gin"
)

// V1Handler 引数を返さないので空
type V1Handler struct{}

func (v V1Handler) GetApiV1Host(c *gin.Context) {
	storageDirEnv := os.Getenv("STORAGE_DIR")
	storageDir := models.GetInfoOptions{
		StorageDir: storageDirEnv,
	}
	getInfo, err := host.GetInfo(storageDir)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	hostInfo := backendHost.GetHostInfo(getInfo)
	c.JSON(http.StatusOK, hostInfo)
}

func (v V1Handler) GetApiV1Vms(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

// PostApiV1Vms Vm作成時にフロントから情報を受取りステータスを返す
func (v V1Handler) PostApiV1Vms(c *gin.Context) {
	var getJsonData models.PostApiV1VmsJSONRequestBody
	if err := c.ShouldBindJSON(&getJsonData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	vm, err := vms.CreateVm(getJsonData)
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
