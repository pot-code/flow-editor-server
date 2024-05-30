// Package flow provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package flow

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

const (
	BearerScopes = "bearer.Scopes"
)

// FlowListItemRes defines model for FlowListItemRes.
type FlowListItemRes struct {
	CreatedAt string `json:"created_at"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
}

// GetFlowJSONBody defines parameters for GetFlow.
type GetFlowJSONBody = []FlowListItemRes

// PostFlowJSONBody defines parameters for PostFlow.
type PostFlowJSONBody struct {
	Edges *string `json:"edges"`
	Nodes *string `json:"nodes"`
	Title string  `json:"title"`
}

// PutFlowJSONBody defines parameters for PutFlow.
type PutFlowJSONBody struct {
	Edges string `json:"edges"`
	Id    string `json:"id"`
	Nodes string `json:"nodes"`
	Title string `json:"title"`
}

// GetFlowIdJSONBody defines parameters for GetFlowId.
type GetFlowIdJSONBody struct {
	CreatedAt string `json:"created_at"`
	Edges     string `json:"edges"`
	Id        int    `json:"id"`
	Nodes     string `json:"nodes"`
	Title     string `json:"title"`
}

// GetFlowJSONRequestBody defines body for GetFlow for application/json ContentType.
type GetFlowJSONRequestBody = GetFlowJSONBody

// PostFlowJSONRequestBody defines body for PostFlow for application/json ContentType.
type PostFlowJSONRequestBody PostFlowJSONBody

// PutFlowJSONRequestBody defines body for PutFlow for application/json ContentType.
type PutFlowJSONRequestBody PutFlowJSONBody

// GetFlowIdJSONRequestBody defines body for GetFlowId for application/json ContentType.
type GetFlowIdJSONRequestBody GetFlowIdJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// 获取 flow 列表
	// (GET /flow)
	GetFlow(ctx echo.Context) error
	// 创建 flow
	// (POST /flow)
	PostFlow(ctx echo.Context) error
	// 更新 flow
	// (PUT /flow)
	PutFlow(ctx echo.Context) error
	// 删除 flow
	// (DELETE /flow/{id})
	DeleteFlowId(ctx echo.Context, id string) error
	// 获取 flow
	// (GET /flow/{id})
	GetFlowId(ctx echo.Context, id string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetFlow converts echo context to params.
func (w *ServerInterfaceWrapper) GetFlow(ctx echo.Context) error {
	var err error

	ctx.Set(BearerScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetFlow(ctx)
	return err
}

// PostFlow converts echo context to params.
func (w *ServerInterfaceWrapper) PostFlow(ctx echo.Context) error {
	var err error

	ctx.Set(BearerScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostFlow(ctx)
	return err
}

// PutFlow converts echo context to params.
func (w *ServerInterfaceWrapper) PutFlow(ctx echo.Context) error {
	var err error

	ctx.Set(BearerScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PutFlow(ctx)
	return err
}

// DeleteFlowId converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteFlowId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set(BearerScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteFlowId(ctx, id)
	return err
}

// GetFlowId converts echo context to params.
func (w *ServerInterfaceWrapper) GetFlowId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set(BearerScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetFlowId(ctx, id)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/flow", wrapper.GetFlow)
	router.POST(baseURL+"/flow", wrapper.PostFlow)
	router.PUT(baseURL+"/flow", wrapper.PutFlow)
	router.DELETE(baseURL+"/flow/:id", wrapper.DeleteFlowId)
	router.GET(baseURL+"/flow/:id", wrapper.GetFlowId)

}
