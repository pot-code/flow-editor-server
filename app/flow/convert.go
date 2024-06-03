package flow

import (
	"time"
)

// goverter:converter
// goverter:output:file convert_impl.go
// goverter:output:package flow-editor-server/app/flow
// goverter:extend TimeToString
// goverter:extend UintToInt
type Converter interface {
	ConvertFlowModels(s []FlowModel) []FlowListObjectOutput
	// goverter:map Model.ID Id
	// goverter:map Model.CreatedAt CreatedAt
	ConvertFlowModel(s FlowModel) FlowDetailOutput
	// goverter:map Model.ID Id
	// goverter:map Model.CreatedAt CreatedAt
	ConverterFlowModel(s FlowModel) FlowListObjectOutput
}

func TimeToString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func UintToInt(t uint) int {
	return int(t)
}
