package flow

import "time"

//go:generate go run github.com/jmattheis/goverter/cmd/goverter gen .

// goverter:converter
// goverter:output:file convert_impl.go
// goverter:output:package openapi-go-demo/app/flow
// goverter:extend TimeToString
type Converter interface {
	ConvertPostFlowJSONRequestBody(s PostFlowJSONRequestBody) CreateFlowData
	ConvertSliceFlowListItem(s []FlowListItem) GetFlowJSONBody
	ConvertFlowDetail(s FlowDetail) GetFlowIdJSONBody
	ConvertPutFlowJSONRequestBody(s PutFlowIdJSONRequestBody) UpdateFlowData
}

func TimeToString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
