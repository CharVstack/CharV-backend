package delivery

import (
	"fmt"
	"github.com/CharVstack/CharV-backend/openapi/v1"
	"github.com/CharVstack/CharV-backend/openapi/v1/usecase/vms"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (v V1Handler) GetApiV1Vms(c *gin.Context) {
	getVmTestData, getVmsInfo := vms.GetVmsInfo()
	fmt.Println("getVmsInfo: ", getVmsInfo)
	//TODO implement me
	c.JSON(http.StatusOK, gin.H{
		"message": getVmTestData,
		"vms":     getVmsInfo,
	})
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
