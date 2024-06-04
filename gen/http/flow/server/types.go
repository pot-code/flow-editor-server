// Code generated by goa v3.16.2, DO NOT EDIT.
//
// flow HTTP server types
//
// Command:
// $ goa gen flow-editor-server/design

package server

import (
	flow "flow-editor-server/gen/flow"

	goa "goa.design/goa/v3/pkg"
)

// CreateFlowRequestBody is the type of the "flow" service "createFlow"
// endpoint HTTP request body.
type CreateFlowRequestBody struct {
	// flow 标题
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// flow 节点
	Nodes *string `form:"nodes,omitempty" json:"nodes,omitempty" xml:"nodes,omitempty"`
	// flow 边
	Edges *string `form:"edges,omitempty" json:"edges,omitempty" xml:"edges,omitempty"`
}

// UpdateFlowRequestBody is the type of the "flow" service "updateFlow"
// endpoint HTTP request body.
type UpdateFlowRequestBody struct {
	// flow 标题
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// flow 节点
	Nodes *string `form:"nodes,omitempty" json:"nodes,omitempty" xml:"nodes,omitempty"`
	// flow 边
	Edges *string `form:"edges,omitempty" json:"edges,omitempty" xml:"edges,omitempty"`
}

// GetFlowListResponseBody is the type of the "flow" service "getFlowList"
// endpoint HTTP response body.
type GetFlowListResponseBody []*FlowListItemResponse

// GetFlowResponseBody is the type of the "flow" service "getFlow" endpoint
// HTTP response body.
type GetFlowResponseBody struct {
	// flow id
	ID int `form:"id" json:"id" xml:"id"`
	// flow 标题
	Title string `form:"title" json:"title" xml:"title"`
	// flow 节点
	Nodes *string `form:"nodes,omitempty" json:"nodes,omitempty" xml:"nodes,omitempty"`
	// flow 边
	Edges *string `form:"edges,omitempty" json:"edges,omitempty" xml:"edges,omitempty"`
	// flow 创建时间
	CreatedAt string `form:"created_at" json:"created_at" xml:"created_at"`
}

// CreateFlowResponseBody is the type of the "flow" service "createFlow"
// endpoint HTTP response body.
type CreateFlowResponseBody struct {
	// flow id
	ID int `form:"id" json:"id" xml:"id"`
	// flow 标题
	Title string `form:"title" json:"title" xml:"title"`
	// flow 节点
	Nodes *string `form:"nodes,omitempty" json:"nodes,omitempty" xml:"nodes,omitempty"`
	// flow 边
	Edges *string `form:"edges,omitempty" json:"edges,omitempty" xml:"edges,omitempty"`
	// flow 创建时间
	CreatedAt string `form:"created_at" json:"created_at" xml:"created_at"`
}

// UpdateFlowResponseBody is the type of the "flow" service "updateFlow"
// endpoint HTTP response body.
type UpdateFlowResponseBody struct {
	// flow id
	ID int `form:"id" json:"id" xml:"id"`
	// flow 标题
	Title string `form:"title" json:"title" xml:"title"`
	// flow 节点
	Nodes *string `form:"nodes,omitempty" json:"nodes,omitempty" xml:"nodes,omitempty"`
	// flow 边
	Edges *string `form:"edges,omitempty" json:"edges,omitempty" xml:"edges,omitempty"`
	// flow 创建时间
	CreatedAt string `form:"created_at" json:"created_at" xml:"created_at"`
}

// FlowListItemResponse is used to define fields on response body types.
type FlowListItemResponse struct {
	// flow id
	ID int `form:"id" json:"id" xml:"id"`
	// flow 标题
	Title string `form:"title" json:"title" xml:"title"`
	// flow 创建时间
	CreatedAt string `form:"created_at" json:"created_at" xml:"created_at"`
}

// NewGetFlowListResponseBody builds the HTTP response body from the result of
// the "getFlowList" endpoint of the "flow" service.
func NewGetFlowListResponseBody(res []*flow.FlowListItem) GetFlowListResponseBody {
	body := make([]*FlowListItemResponse, len(res))
	for i, val := range res {
		body[i] = marshalFlowFlowListItemToFlowListItemResponse(val)
	}
	return body
}

// NewGetFlowResponseBody builds the HTTP response body from the result of the
// "getFlow" endpoint of the "flow" service.
func NewGetFlowResponseBody(res *flow.FlowDetail) *GetFlowResponseBody {
	body := &GetFlowResponseBody{
		ID:        res.ID,
		Title:     res.Title,
		Nodes:     res.Nodes,
		Edges:     res.Edges,
		CreatedAt: res.CreatedAt,
	}
	return body
}

// NewCreateFlowResponseBody builds the HTTP response body from the result of
// the "createFlow" endpoint of the "flow" service.
func NewCreateFlowResponseBody(res *flow.FlowDetail) *CreateFlowResponseBody {
	body := &CreateFlowResponseBody{
		ID:        res.ID,
		Title:     res.Title,
		Nodes:     res.Nodes,
		Edges:     res.Edges,
		CreatedAt: res.CreatedAt,
	}
	return body
}

// NewUpdateFlowResponseBody builds the HTTP response body from the result of
// the "updateFlow" endpoint of the "flow" service.
func NewUpdateFlowResponseBody(res *flow.FlowDetail) *UpdateFlowResponseBody {
	body := &UpdateFlowResponseBody{
		ID:        res.ID,
		Title:     res.Title,
		Nodes:     res.Nodes,
		Edges:     res.Edges,
		CreatedAt: res.CreatedAt,
	}
	return body
}

// NewCreateFlowData builds a flow service createFlow endpoint payload.
func NewCreateFlowData(body *CreateFlowRequestBody) *flow.CreateFlowData {
	v := &flow.CreateFlowData{
		Title: *body.Title,
		Nodes: body.Nodes,
		Edges: body.Edges,
	}

	return v
}

// NewUpdateFlowPayload builds a flow service updateFlow endpoint payload.
func NewUpdateFlowPayload(body *UpdateFlowRequestBody, id string) *flow.UpdateFlowPayload {
	v := &flow.UpdateFlowData{
		Title: body.Title,
		Nodes: body.Nodes,
		Edges: body.Edges,
	}
	res := &flow.UpdateFlowPayload{
		Data: v,
	}
	res.ID = &id

	return res
}

// ValidateCreateFlowRequestBody runs the validations defined on
// CreateFlowRequestBody
func ValidateCreateFlowRequestBody(body *CreateFlowRequestBody) (err error) {
	if body.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "body"))
	}
	return
}
