// Package adapters provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
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
	GetApiV1VmsVmId(c *gin.Context, vmId string)
	// Update a VM
	// (PATCH /api/v1/vms/{vmId})
	PatchApiV1VmsVmId(c *gin.Context, vmId string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
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
	var vmId string

	err = runtime.BindStyledParameter("simple", false, "vmId", c.Param("vmId"), &vmId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter vmId: %s", err)})
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
	var vmId string

	err = runtime.BindStyledParameter("simple", false, "vmId", c.Param("vmId"), &vmId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter vmId: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.PatchApiV1VmsVmId(c, vmId)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL     string
	Middlewares []MiddlewareFunc
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router *gin.Engine, si ServerInterface) *gin.Engine {
	return RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router *gin.Engine, si ServerInterface, options GinServerOptions) *gin.Engine {
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
	}

	router.GET(options.BaseURL+"/api/v1/host", wrapper.GetApiV1Host)

	router.GET(options.BaseURL+"/api/v1/vms", wrapper.GetApiV1Vms)

	router.POST(options.BaseURL+"/api/v1/vms", wrapper.PostApiV1Vms)

	router.GET(options.BaseURL+"/api/v1/vms/:vmId", wrapper.GetApiV1VmsVmId)

	router.PATCH(options.BaseURL+"/api/v1/vms/:vmId", wrapper.PatchApiV1VmsVmId)

	return router
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"UB+Mx/8EAAD//28nNdCpFgAA",
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
