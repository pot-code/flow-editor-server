// Code generated by goa v3.16.2, DO NOT EDIT.
//
// flow service
//
// Command:
// $ goa gen flow-editor-server/design

package flow

import (
	"context"
)

// Flow 服务
type Service interface {
	// 列出当前用户拥有的 flow
	GetFlowList(context.Context) (res []*FlowListItem, err error)
	// 根据 flow id 获取 flow 详情
	GetFlow(context.Context, int) (res *FlowDetail, err error)
	// 创建 flow
	CreateFlow(context.Context, *CreateFlowData) (res *FlowDetail, err error)
	// 更新 flow
	UpdateFlow(context.Context, *UpdateFlowPayload) (res *FlowDetail, err error)
	// 删除 flow
	DeleteFlow(context.Context, int) (err error)
}

// APIName is the name of the API as defined in the design.
const APIName = "flow-editor-server"

// APIVersion is the version of the API as defined in the design.
const APIVersion = "0.0.1"

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "flow"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [5]string{"getFlowList", "getFlow", "createFlow", "updateFlow", "deleteFlow"}

// CreateFlowData is the payload type of the flow service createFlow method.
type CreateFlowData struct {
	Title string
	Nodes *string
	Edges *string
}

// FlowDetail is the result type of the flow service getFlow method.
type FlowDetail struct {
	ID        int
	Title     string
	Nodes     *string
	Edges     *string
	CreatedAt string
}

type FlowListItem struct {
	ID        int
	Title     string
	CreatedAt string
}

type UpdateFlowData struct {
	Title *string
	Nodes *string
	Edges *string
}

// UpdateFlowPayload is the payload type of the flow service updateFlow method.
type UpdateFlowPayload struct {
	Data *UpdateFlowData
	ID   *int
}