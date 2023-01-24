package controller

import (
	"fmt"
	"net/http"

	"github.com/CharVstack/CharV-backend/infrastructure/system"
	"github.com/CharVstack/CharV-backend/infrastructure/vnc"
	"github.com/CharVstack/CharV-backend/middleware"
	"github.com/gamoutatsumi/go-vncproxy"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/net/websocket"
)

type vncHandler struct {
	logger   vncproxy.Logger
	logLevel uint32
	path     system.Paths
}

type vncLogger struct {
	logger *zap.Logger
}

func (vh *vncHandler) Handler(c *gin.Context) {
	id := c.Param("vmId")

	isExist := vh.path.Search(system.VNC, id+".sock")
	if !isExist {
		middleware.GenericErrorHandler(c, fmt.Errorf("%s is not found", id), http.StatusNotFound)
		return
	}

	proxy := vnc.NewVNCProxy(id, vh.logger, vh.path.VNC, vh.logLevel)

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

func NewVNCHandler(logger *zap.Logger, path system.Paths, isProduction bool) *vncHandler {
	vncLogger := newVNCLogger(logger)
	var logLevel uint32
	if isProduction {
		logLevel = vncproxy.InfoLevel
	} else {
		logLevel = vncproxy.DebugLevel
	}
	return &vncHandler{
		logger:   vncLogger,
		logLevel: logLevel,
		path:     path,
	}
}
