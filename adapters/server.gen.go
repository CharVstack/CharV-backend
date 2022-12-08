// Package adapters provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package adapters

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get a host
	// (GET /api/v1/host)
	GetApiV1Host(c *gin.Context)
	// Get all VMs list
	// (GET /api/v1/vms)
	GetApiV1Vms(c *gin.Context)
	// Create a new VM
	// (POST /api/v1/vms)
	PostApiV1Vms(c *gin.Context)
	// Get a VM
	// (GET /api/v1/vms/{vmId})
	GetApiV1VmsVmId(c *gin.Context, vmId openapi_types.UUID)
	// Update a VM
	// (PATCH /api/v1/vms/{vmId})
	PatchApiV1VmsVmId(c *gin.Context, vmId openapi_types.UUID)
	// Get Power State
	// (GET /api/v1/vms/{vmId}/power)
	GetApiV1VmsVmIdPower(c *gin.Context, vmId openapi_types.UUID)
	// Change Power Status
	// (POST /api/v1/vms/{vmId}/power)
	PostApiV1VmsVmIdPower(c *gin.Context, vmId openapi_types.UUID)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// GetApiV1Host operation middleware
func (siw *ServerInterfaceWrapper) GetApiV1Host(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetApiV1Host(c)
}

// GetApiV1Vms operation middleware
func (siw *ServerInterfaceWrapper) GetApiV1Vms(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetApiV1Vms(c)
}

// PostApiV1Vms operation middleware
func (siw *ServerInterfaceWrapper) PostApiV1Vms(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.PostApiV1Vms(c)
}

// GetApiV1VmsVmId operation middleware
func (siw *ServerInterfaceWrapper) GetApiV1VmsVmId(c *gin.Context) {

	var err error

	// ------------- Path parameter "vmId" -------------
	var vmId openapi_types.UUID

	err = runtime.BindStyledParameter("simple", false, "vmId", c.Param("vmId"), &vmId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter vmId: %s", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetApiV1VmsVmId(c, vmId)
}

// PatchApiV1VmsVmId operation middleware
func (siw *ServerInterfaceWrapper) PatchApiV1VmsVmId(c *gin.Context) {

	var err error

	// ------------- Path parameter "vmId" -------------
	var vmId openapi_types.UUID

	err = runtime.BindStyledParameter("simple", false, "vmId", c.Param("vmId"), &vmId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter vmId: %s", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.PatchApiV1VmsVmId(c, vmId)
}

// GetApiV1VmsVmIdPower operation middleware
func (siw *ServerInterfaceWrapper) GetApiV1VmsVmIdPower(c *gin.Context) {

	var err error

	// ------------- Path parameter "vmId" -------------
	var vmId openapi_types.UUID

	err = runtime.BindStyledParameter("simple", false, "vmId", c.Param("vmId"), &vmId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter vmId: %s", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetApiV1VmsVmIdPower(c, vmId)
}

// PostApiV1VmsVmIdPower operation middleware
func (siw *ServerInterfaceWrapper) PostApiV1VmsVmIdPower(c *gin.Context) {

	var err error

	// ------------- Path parameter "vmId" -------------
	var vmId openapi_types.UUID

	err = runtime.BindStyledParameter("simple", false, "vmId", c.Param("vmId"), &vmId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter vmId: %s", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.PostApiV1VmsVmIdPower(c, vmId)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router *gin.Engine, si ServerInterface) *gin.Engine {
	return RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router *gin.Engine, si ServerInterface, options GinServerOptions) *gin.Engine {

	errorHandler := options.ErrorHandler

	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.GET(options.BaseURL+"/api/v1/host", wrapper.GetApiV1Host)

	router.GET(options.BaseURL+"/api/v1/vms", wrapper.GetApiV1Vms)

	router.POST(options.BaseURL+"/api/v1/vms", wrapper.PostApiV1Vms)

	router.GET(options.BaseURL+"/api/v1/vms/:vmId", wrapper.GetApiV1VmsVmId)

	router.PATCH(options.BaseURL+"/api/v1/vms/:vmId", wrapper.PatchApiV1VmsVmId)

	router.GET(options.BaseURL+"/api/v1/vms/:vmId/power", wrapper.GetApiV1VmsVmIdPower)

	router.POST(options.BaseURL+"/api/v1/vms/:vmId/power", wrapper.PostApiV1VmsVmIdPower)

	return router
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
