package router

import (
	"os"
	"time"

	"github.com/CharVstack/CharV-backend/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewRouter(logger *zap.Logger) *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())

	r.Use(middleware.Logger(logger))

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			os.Getenv("ORIGIN_URI"),
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"OPTIONS",
			"DELETE",
		},
		AllowHeaders: []string{
			"Content-Type",
		},
		AllowCredentials: false,
		MaxAge:           24 * time.Hour,
	}))

	return r
}
