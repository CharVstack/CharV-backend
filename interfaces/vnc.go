package interfaces

import (
	"fmt"
	"github.com/CharVstack/CharV-backend/infrastructure/vnc"
	"github.com/CharVstack/CharV-backend/usecase/models"
	"net/http"

	"github.com/CharVstack/CharV-backend/middleware"
	"github.com/gamoutatsumi/go-vncproxy"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/net/websocket"
)

type vncHandler struct {
	logger     vncproxy.Logger
	socketsDir string
	logLevel   uint32
	da         models.VmDataAccess
}

type vncLogger struct {
	logger *zap.Logger
}

func (vh *vncHandler) Handler(c *gin.Context) {
	vmId := c.Param("vmId")
	vms, err := vh.da.Browse()
	if err != nil {
		middleware.GenericErrorHandler(c, err, http.StatusInternalServerError)
		return
	}
	var proxy *vncproxy.Proxy
	for _, vm := range vms {
		if vmId == vm.ID.String() {
			proxy = vnc.NewVNCProxy(vmId, vh.logger, vh.socketsDir, vh.logLevel)
			break
		}
	}
	if proxy == nil {
		middleware.GenericErrorHandler(c, fmt.Errorf("%s is not found", vmId), http.StatusNotFound)
		return
	}
	h := websocket.Handler(proxy.ServeWS)
	h.ServeHTTP(c.Writer, c.Request)
}

func (l *vncLogger) Infof(format string, v ...interface{}) {
	l.logger.Info(fmt.Sprintf(format, fmt.Sprint(v...)), zap.String("provider", "vncproxy"))
}

func (l *vncLogger) Info(msg string) {
	l.logger.Info(msg, zap.String("provider", "vncproxy"))
}

func (l *vncLogger) Debugf(format string, v ...interface{}) {
	l.logger.Debug(fmt.Sprintf(format, fmt.Sprint(v...)), zap.String("provider", "vncproxy"))
}

func (l *vncLogger) Debug(msg string) {
	l.logger.Debug(msg, zap.String("provider", "vncproxy"))
}

func newVNCLogger(logger *zap.Logger) *vncLogger {
	return &vncLogger{
		logger: logger,
	}
}

func NewVNCHandler(logger *zap.Logger, socketsDir string, isProduction bool, da *models.VmDataAccess) *vncHandler {
	vncLogger := newVNCLogger(logger)
	var logLevel uint32
	if isProduction {
		logLevel = vncproxy.InfoLevel
	} else {
		logLevel = vncproxy.DebugLevel
	}
	return &vncHandler{
		logger:     vncLogger,
		socketsDir: socketsDir,
		logLevel:   logLevel,
		da:         *da,
	}
}
