package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(production bool) (*zap.Logger, error) {
	if production {
		config := zap.NewProductionConfig()
		config.Encoding = "console"
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

		gin.SetMode(gin.ReleaseMode)

		return config.Build()
	} else {
		return zap.NewDevelopmentConfig().Build()
	}
}
