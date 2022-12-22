package handler

import (
	"fmt"
	"net/http"

	"github.com/CharVstack/CharV-backend/middleware"
	"github.com/CharVstack/CharV-backend/pkg/qemu"
	"github.com/gamoutatsumi/go-vncproxy"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/websocket"
)

func VNCHandler(c *gin.Context, vmId string) {
	vms, err := qemu.GetAllVmInfo()
	if err != nil {
		middleware.GenericErrorHandler(c, err, http.StatusInternalServerError)
		return
	}
	var proxy *vncproxy.Proxy
	for _, vm := range vms {
		if vmId == vm.Metadata.Id.String() {
			proxy = qemu.NewVNCProxy(vmId)
			break
		}
	}
	if proxy == nil {
		middleware.GenericErrorHandler(c, fmt.Errorf("%s is not found.", vmId), http.StatusNotFound)
		return
	}
	h := websocket.Handler(proxy.ServeWS)
	h.ServeHTTP(c.Writer, c.Request)
}
