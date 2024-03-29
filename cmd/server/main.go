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
	"log"
	"os"

	"github.com/CharVstack/CharV-backend/api"
	"github.com/CharVstack/CharV-backend/infrastructure/disk"
	"github.com/CharVstack/CharV-backend/infrastructure/file"
	"github.com/CharVstack/CharV-backend/infrastructure/system"
	"github.com/CharVstack/CharV-backend/infrastructure/utils"
	"github.com/CharVstack/CharV-backend/interfaces/controller"
	"github.com/CharVstack/CharV-backend/interfaces/middleware"
	"github.com/CharVstack/CharV-backend/interfaces/router"
	"github.com/CharVstack/CharV-backend/usecase/host"
	"github.com/CharVstack/CharV-backend/usecase/vm/qemu"
	oapiMiddleware "github.com/deepmap/oapi-codegen/pkg/gin-middleware"
)

var production bool

func main() {
	logger, err := router.NewLogger(production)
	if err != nil {
		log.Fatal(err)
	}

	r := router.NewRouter(logger)

	sys := system.Paths{
		Images:       os.Getenv("IMAGES_DIR"),
		Guests:       os.Getenv("GUESTS_DIR"),
		StoragePools: os.Getenv("STORAGE_POOLS_DIR"),
		QMP:          os.Getenv("QMP_DIR"),
		VNC:          os.Getenv("VNC_DIR"),
	}

	// ToDo: DI ライブラリの導入を検討する

	d1 := file.NewVmDataAccess(sys)
	d2 := file.NewStorageAccess(sys)

	stat := utils.NewHostStatAccess(&d2)
	qcow2 := disk.NewQCOW2Disk(&d2, sys)

	u1 := qemu.NewQemuUseCase(&d1, &qcow2, sys)
	u2 := host.NewHostUseCase(&stat)

	v1Handler := controller.NewV1Handler(&u1, &u2)

	ginServerOpts := api.GinServerOptions{
		ErrorHandler: middleware.GenericErrorHandler,
	}

	vncHandler := controller.NewVNCHandler(logger, sys, production)
	r.GET("/ws/vnc/:vmId", vncHandler.Handler)

	route := api.RegisterHandlersWithOptions(r, v1Handler, ginServerOpts)

	swagger, err := api.GetSwagger()
	if err != nil {
		log.Fatal(err.Error())
	}
	validatorOpts := oapiMiddleware.Options{
		ErrorHandler: middleware.ValidationErrorHandler,
	}

	oasRouter := route.Group("/api")
	oasRouter.Use(oapiMiddleware.OapiRequestValidatorWithOptions(swagger, &validatorOpts))

	if err := route.Run(":4010"); err != nil {
		log.Fatal(err.Error())
	}
}
