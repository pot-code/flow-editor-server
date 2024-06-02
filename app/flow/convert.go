package flow

import "time"

//go:generate go run github.com/jmattheis/goverter/cmd/goverter gen .

// goverter:converter
// goverter:output:file convert_impl.go
// goverter:output:package openapi-go-demo/app/flow
// goverter:extend TimeToString
// goverter:extend UintToInt
type Converter interface {
	ConvertPostFlowJSONRequestBody(s PostFlowJSONRequestBody) CreateFlowData
	ConvertSliceFlowListItem(s []FlowListItem) []FlowListObject
	ConvertPutFlowJSONRequestBody(s PutFlowIdJSONRequestBody) UpdateFlowData
	// goverter:map Model.ID Id
	// goverter:map Model.CreatedAt CreatedAt
	ConvertFlowModel(s FlowModel) FlowDetailObject
}

func TimeToString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func UintToInt(t uint) int {
	return int(t)
}
