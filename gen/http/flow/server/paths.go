// Code generated by goa v3.17.1, DO NOT EDIT.
//
// HTTP request path constructors for the flow service.
//
// Command:
// $ goa gen flow-editor-server/design

package server

import (
	"fmt"
)

// GetFlowListFlowPath returns the URL path to the flow service getFlowList HTTP endpoint.
func GetFlowListFlowPath() string {
	return "/flow"
}

// GetFlowFlowPath returns the URL path to the flow service getFlow HTTP endpoint.
func GetFlowFlowPath(id string) string {
	return fmt.Sprintf("/flow/%v", id)
}

// CreateFlowFlowPath returns the URL path to the flow service createFlow HTTP endpoint.
func CreateFlowFlowPath() string {
	return "/flow"
}

// UpdateFlowFlowPath returns the URL path to the flow service updateFlow HTTP endpoint.
func UpdateFlowFlowPath(id string) string {
	return fmt.Sprintf("/flow/%v", id)
}

// DeleteFlowFlowPath returns the URL path to the flow service deleteFlow HTTP endpoint.
func DeleteFlowFlowPath(id string) string {
	return fmt.Sprintf("/flow/%v", id)
}

// CopyFlowFlowPath returns the URL path to the flow service copyFlow HTTP endpoint.
func CopyFlowFlowPath(id string) string {
	return fmt.Sprintf("/flow/%v/copy", id)
}
