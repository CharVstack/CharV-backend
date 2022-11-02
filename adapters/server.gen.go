// Package adapters provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.2 DO NOT EDIT.
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
	PostApiV1VmsVmIdPowerAction(c *gin.Context, vmId openapi_types.UUID, params PostApiV1VmsVmIdPowerActionParams)
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

// PostApiV1VmsVmIdPowerAction operation middleware
func (siw *ServerInterfaceWrapper) PostApiV1VmsVmIdPowerAction(c *gin.Context) {

	var err error

	// ------------- Path parameter "vmId" -------------
	var vmId openapi_types.UUID

	err = runtime.BindStyledParameter("simple", false, "vmId", c.Param("vmId"), &vmId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter vmId: %s", err), http.StatusBadRequest)
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params PostApiV1VmsVmIdPowerActionParams

	// ------------- Optional query parameter "action" -------------

	err = runtime.BindQueryParameter("form", true, false, "action", c.Request.URL.Query(), &params.Action)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter action: %s", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.PostApiV1VmsVmIdPowerAction(c, vmId, params)
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

	router.POST(options.BaseURL+"/api/v1/vms/:vmId/power", wrapper.PostApiV1VmsVmIdPowerAction)

	return router
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xZbY/bxhH+K8K2H3kniqJeTt8c38U9oHKEGFY/BAdhRa4k+kguzV3qrAgCrBOMGLWL",
	"GC7qoHDQNkXhBnVhpEjQpk3TP0Of6v6LYl/4IpHSSfZdiyT+IEAklzvzzDwzOzMcAwM7HnaRSwlojIGP",
	"iIddgvjFNUSv2Ha7SX5qEaqp6vvyIXtmYJcil7K/0PNsy4DUwm7xFsEuu4fuQMezxTby/06JXQyhHfAN",
	"TEghaHwwBiYaWoZYaVrkOHUPNMQdBXiQDkADFIfQL9pWl/2Glk+LlgP7iBSDbuDSQNN2VX33toFPNKAA",
	"OvLYBuJycjRRgIMc7I9AQ1P1OruiUOjAEFidIfKJxZQHwxJQgGWCBtCRCsuGXt/p6fXyjq5pcKdb0dQd",
	"ra6pmlozUE/XwEQBLnSYLKEGUAChkAYENAA0qDVEQAFDwwtAQ58cTSYTBRBjgBwu2fOxh3xqSfxSH4si",
	"h9/4sY96oAF+VEy8VBQvk2LbYaIlTOj7cAQ4SEJgn1tYPiLUt9w+YHJ9dDuwfGSCxgdC1lG8Ae7eQgYF",
	"Crizs9p3TKnvp8fS/mGWMhExfMujYv+ze5+H02fh9EW7+fLru6+e/ZFtcQ3Rn+DLiosx4Aqx/QIemqoC",
	"POQbfGtVmIY97vkI8YcUU2jzfwFhDpZ/OgsvEYp92EcdD2NbuFLaQXIkdlt8HTP5SsRkLqhDrA9RIkRe",
	"SZ9JAibE24Dx64jOrHwB1M7xazh7Gp7+PZzdD6cv5rN7Z7/7i3Rsu/nOqN08NC/LuT+onJfmxJuzQaS9",
	"i+cCc3c4fX726P6rj74Kp78OTx+0m0ukaOET5F8uMwwbQbfjMUEd3OuBBvUDJKzL9G+997OD9w/2O++9",
	"+y64lGBrOxzkodvDl2XncPriP0+/mf/jUWLdFqTG4KZnQoreht4PLvTmT7+aP/kinP4qPH0YTn+bDrsW",
	"JvSqjyBF19FJu/mWFNuTwkOuKQ7z7xgrnnzx6tnHL7/9dH7/UQ43Yq25oldFtbT6eL/auinfVJYwRhVW",
	"rLjlUtRHPkMVV09j0MO+AymjAw66Nkq86wZOl61ewim3TfZgmC1qs1eYtudV3ibqwcCmaQ31lEIVJu/O",
	"DoV9VsiBAauRGMH2EyYvuZLzesPeYp8tznQXy55ki1KoItF5yAjFnm31BxwP5/Wob5VHH1aPrcGeavO9",
	"96WGS3rLMBwD5AZOLFcBhuljByggcI9dfOKmOBUxL4rZcfaBuJHsGW2iyGDNbrYEnj9VIu2kpLQxhJKb",
	"WILax70Tstf1oVfvcUG84s1YQnYE6/zGiJU0B+tWNkUiyjYFGzLkhnirhbF9LlEMTnem1LK0lME45s0M",
	"Zg893amX9kYYuTUuTaJZmwHC2Wfh7Pfh7E8r8oBopVKBHlgureqJTqm8ILutzRaLhmzztZ03yDpCMSlT",
	"EZiWdk3ZXJpti1QkjFQuVXRVLe1psSVK1YpWrdeq1ZoWAa7X9b1aTavo+jKsSnm3otdq1VJNL9dK5T19",
	"RS5rpo7HRV8tHJY54W0tGTywTHBeRPMl6Y0XDCU12YygvmeXTkaqqdl3+oQLSofLOpY+nD+chtM/yMvZ",
	"n8PZP8PTr8PZJ+zP7HnSpy6zVxz/4y1SYFQgJEkw7vEPfB/7uQk13f1vwejNX1hyCoe1IDa9o0QXY0l5",
	"LG3wLfj9ZvOQkqpWK2W1WtE1tb4AnfG8qu6VKnqluoLtbSfLjJffvJjPvgxnvwlP/xbOvgxPH7/67HPW",
	"nLJE9lE4e57hQaqSXXvAy2WpkjSv/kkXqOtPE7kuLkQ3IhziPFOSviWpVded6kN5Em7GHSeJXSfKd0Nx",
	"JJlxxZIlUNvZhjd5DcR3qWHQ1F1VX9tKLpB26EjKJsOCbLmyPMmIHdbFmD2LOLFQiKUHHEpydR0o4MbN",
	"G62D6/sH++dXZ8uyI0EL/k103yyrn3j9Qc8J3L6puRUg5sSWhB5X9gPotwmFxvHOldYhM1/sodKuyhBj",
	"D7nQs0ADlHfVXVXmGG6xIvSs4rBUHMjyr4/omsNCHATh6eOzj5+c/esTMbICXIDP++BDpvQ1RK94Vrsk",
	"y6uF7zuaqq4K7HhdMWfMzXuvwHEgIyaTUICFgSzfFrPaRIlBDUVNmYspM1/fFFXbIa8LKv/LVg402y60",
	"m6RgWwsAZQR40lWZGdfp46hzzQXQwmQRwe0AEfoONkcXMtdYzBq5kR7H9Zr2f93RsDLNb5ud0yk5by4w",
	"eR0Hr5kbLbpYLCrAgotOCu3msocXCVwcD51Dc7Kax3cfhHfvcR5vF51th7SdQ/N1uZw3M80L0iw+ln18",
	"6CCKfPFByGI4ZFUleTOMNIucJ0fSMWnOq7OPeI4zBnmBEtspGgLmhwt7PWOpy4mZkqrpMXbWsGDSIdRH",
	"0ImjRvufRs0FRcT6EfsiXcS6XMbkRkSRH7Ur4yI18//3z/86v/cgjovzIoKf0Vs7O/HMxp86XtOq6z4L",
	"ZSOQLyzc4KXI/ysQVx1YGQfNf/mLl99+uvbkij3EuiHsgnwQtwPEM7xEAaO1id5R8Uco9NkhSwLCOgGR",
	"DhEVYyMvr+w7yjhNz6K7jgtXJWGWkv8Aun2UcktAMnSfxDfGiR/ARImvBuLzdHzNVWSK5dWQ2l4JeYbR",
	"7d+qaSaYTP4bAAD//z9fkR6AIwAA",
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
