/*
 * CharVstack-API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"errors"
	"flag"
	"io/fs"
	"log"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/CharVstack/CharV-backend/adapters"
	"github.com/CharVstack/CharV-backend/handler"
	"github.com/CharVstack/CharV-backend/middleware"
	oapiMiddleware "github.com/deepmap/oapi-codegen/pkg/gin-middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var production bool

func init() {
	var (
		configPath = flag.String("c", "/etc/charv/backend.conf", "backend config file path")
	)
	flag.Parse()

	err := godotenv.Load(*configPath)
	if err != nil {
		log.Fatal(err.Error())
	}

	storageDirEnv := os.Getenv("STORAGE_DIR")
	_, err = os.ReadDir(storageDirEnv)
	if err != nil && errors.Is(err, fs.ErrNotExist) {
		log.Fatal(err.Error())
	}

	socketsDirEnv := os.Getenv("SOCKETS_DIR")
	_, err = os.ReadDir(socketsDirEnv)
	if err != nil && errors.Is(err, fs.ErrNotExist) {
		log.Fatal(err.Error())
	}
}

func main() {
	var logger *zap.Logger
	if production {
		var err error

		config := zap.NewProductionConfig()
		config.Encoding = "console"
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

		logger, err = config.Build()
		if err != nil {
			log.Fatal(err)
		}

		gin.SetMode(gin.ReleaseMode)
	} else {
		var err error
		logger, err = zap.NewDevelopmentConfig().Build()
		if err != nil {
			log.Fatal(err)
		}
	}

	r := gin.New()

	swagger, err := adapters.GetSwagger()
	if err != nil {
		log.Fatal(err.Error())
	}
	validatorOpts := oapiMiddleware.Options{
		ErrorHandler: middleware.ValidationErrorHandler,
	}
	r.Use(oapiMiddleware.OapiRequestValidatorWithOptions(swagger, &validatorOpts))

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
		},
		AllowHeaders: []string{
			"Content-Type",
		},
		AllowCredentials: false,
		MaxAge:           24 * time.Hour,
	}))

	vncHandler := handler.NewVNCHandler(logger)
	r.GET("/ws/vnc/:vmId", vncHandler.Handler)

	opts := handler.ServerConfig{
		StorageDir: os.Getenv("STORAGE_DIR"),
		SocketsDir: os.Getenv("SOCKETS_DIR"),
	}
	v1Handler := handler.NewV1Handler(opts)

	ginServerOpts := adapters.GinServerOptions{
		ErrorHandler: middleware.GenericErrorHandler,
	}
	router := adapters.RegisterHandlersWithOptions(r, v1Handler, ginServerOpts)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err.Error())
	}
}
