package flow

import (
	"flow-editor-server/gen/flow"
	"time"
)

// goverter:converter
// goverter:output:file convert_impl.go
// goverter:output:package flow-editor-server/app/flow
// goverter:extend TimeToString
// goverter:extend UintToInt
type Converter interface {
	FlowModelsToFlowList(s []*FlowModel) []*flow.FlowListItem
	// goverter:map Model.ID ID
	// goverter:map Model.CreatedAt CreatedAt
	FlowModelToFlowDetail(s *FlowModel) *flow.FlowDetail
	// goverter:map Model.ID ID
	// goverter:map Model.CreatedAt CreatedAt
	FlowModelToFlowListItem(s *FlowModel) *flow.FlowListItem
}

func TimeToString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func UintToInt(t uint) int {
	return int(t)
}
