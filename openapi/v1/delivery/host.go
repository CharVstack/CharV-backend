package delivery

import (
	backendHost "github.com/CharVstack/CharV-backend/openapi/v1/usecase/host"
	"net/http"

	"github.com/CharVstack/CharV-lib/host"
	"github.com/gin-gonic/gin"
)

type V1Handler struct{}

func (v V1Handler) GetApiV1Host(c *gin.Context) {
	getInfo := host.GetInfo()
	hostInfo := backendHost.GetHostInfo(getInfo)
	c.JSON(http.StatusOK, gin.H{
		"cpu":           hostInfo.Cpu,
		"mem":           hostInfo.Mem,
		"storage_pools": hostInfo.StoragePools,
	})
}
