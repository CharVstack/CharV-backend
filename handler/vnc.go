package handler

import (
	"fmt"
	"net/http"

	"github.com/CharVstack/CharV-backend/middleware"
	"github.com/CharVstack/CharV-backend/pkg/qemu"
	"github.com/gamoutatsumi/go-vncproxy"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/net/websocket"
)

type vncHandler struct {
	logger     vncproxy.Logger
	socketsDir string
}

type vncLogger struct {
	logger *zap.Logger
}

func (vh *vncHandler) Handler(c *gin.Context) {
	vmId := c.Param("vmId")
	vms, err := qemu.GetAllVmInfo()
	if err != nil {
		middleware.GenericErrorHandler(c, err, http.StatusInternalServerError)
		return
	}
	var proxy *vncproxy.Proxy
	for _, vm := range vms {
		if vmId == vm.Metadata.Id.String() {
			proxy = qemu.NewVNCProxy(vmId, vh.logger, vh.socketsDir)
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

func (l *vncLogger) Infof(format string, v ...interface{}) {
	l.logger.Info(fmt.Sprintf(format, fmt.Sprint(v...)), zap.String("provider", "vncproxy"))
}

func (l *vncLogger) Debugf(format string, v ...interface{}) {
	l.logger.Debug(fmt.Sprintf(format, fmt.Sprint(v...)), zap.String("provider", "vncproxy"))
}

func newVNCLogger(logger *zap.Logger) *vncLogger {
	return &vncLogger{
		logger: logger,
	}
}

func NewVNCHandler(logger *zap.Logger, socketsDir string) *vncHandler {
	vncLogger := newVNCLogger(logger)
	return &vncHandler{
		logger:     vncLogger,
		socketsDir: socketsDir,
	}
}
