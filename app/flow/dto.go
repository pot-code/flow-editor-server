package flow

import "time"

type CreateFlow struct {
	Title string `validate:"required,min=1,max=100"`
	Nodes string
	Edges string
}

type FlowListItem struct {
	Id        int
	Title     string
	CreatedAt time.Time
}

type FlowDetail struct {
	Id        int
	Title     string
	Nodes     string
	Edges     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// goverter:converter
// goverter:output:file dto_gen.go
// goverter:output:package openapi-go-demo/app/flow
// goverter:extend TimeToString
type Converter interface {
	ConvertCreateFlow(s PostFlowJSONRequestBody) CreateFlow
	ConvertListFlow(s []FlowListItem) GetFlowJSONBody
	ConvertFlowDetail(s FlowDetail) GetFlowIdJSONBody
}

func TimeToString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
