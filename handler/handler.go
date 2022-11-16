package handler

import (
	"errors"
	"net/http"
	"os"

	"github.com/CharVstack/CharV-backend/pkg/qemu"
	"github.com/CharVstack/CharV-backend/pkg/util"

	"github.com/CharVstack/CharV-backend/middleware"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"

	"github.com/CharVstack/CharV-backend/domain/models"
	"github.com/CharVstack/CharV-backend/pkg/host"
	"github.com/gin-gonic/gin"
)

// V1Handler 引数を返さないので空
type V1Handler struct{}

func (v V1Handler) GetApiV1Host(c *gin.Context) {
	storageDir := os.Getenv("STORAGE_DIR")
	getHostInfoOpts := host.GetInfoOptions{
		StorageDir: storageDir,
	}
	hostInfo, err := host.GetInfo(getHostInfoOpts)
	if err != nil {
		middleware.GenericErrorHandler(c, err, http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, models.GetHost200Response{
		Data: hostInfo,
	})
}

func (v V1Handler) GetApiV1Vms(c *gin.Context) {
	allVmsInfo, err := qemu.GetAllVmInfo()
	if err != nil {
		middleware.GenericErrorHandler(c, err, http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, models.GetAllVMsList200Response{
		Data: allVmsInfo,
	})
}

// PostApiV1Vms Vm作成時にフロントから情報を受取りステータスを返す
func (v V1Handler) PostApiV1Vms(c *gin.Context) {
	var requestBody models.PostApiV1VmsJSONRequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		middleware.GenericErrorHandler(c, err, http.StatusBadRequest)
		return
	}

	err := util.ExistsSufficientMemory(uint64(requestBody.Memory))
	if err != nil {
		middleware.GenericErrorHandler(c, err, http.StatusInternalServerError)
		return
	}

	vm, err := qemu.CreateVm(requestBody)
	if err != nil {
		middleware.GenericErrorHandler(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, models.PostCreateNewVM200Response{
		Data: vm,
	})
}

func (v V1Handler) GetApiV1VmsVmId(c *gin.Context, vmId openapi_types.UUID) {
	//TODO implement me
	middleware.GenericErrorHandler(c, errors.New("implement me"), http.StatusInternalServerError)
	return
}

func (v V1Handler) PatchApiV1VmsVmId(c *gin.Context, vmId openapi_types.UUID) {
	//TODO implement me
	middleware.GenericErrorHandler(c, errors.New("implement me"), http.StatusInternalServerError)
	return
}

func (v V1Handler) GetApiV1VmsVmIdPower(c *gin.Context, vmId openapi_types.UUID) {
	//TODO implement me
	middleware.GenericErrorHandler(c, errors.New("implement me"), http.StatusInternalServerError)
	return
}

func (v V1Handler) PostApiV1VmsVmIdPowerAction(c *gin.Context, vmId openapi_types.UUID, params models.PostApiV1VmsVmIdPowerActionParams) {
	//TODO implement me
	middleware.GenericErrorHandler(c, errors.New("implement me"), http.StatusInternalServerError)
	return
}
