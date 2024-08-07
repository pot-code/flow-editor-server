// Code generated by goa v3.17.1, DO NOT EDIT.
//
// account service
//
// Command:
// $ goa gen flow-editor-server/design

package account

import (
	"context"
)

// Account service
type Service interface {
	// Get account
	GetAccount(context.Context) (res *AccountInfo, err error)
}

// APIName is the name of the API as defined in the design.
const APIName = "flow-editor-server"

// APIVersion is the version of the API as defined in the design.
const APIVersion = "0.0.1"

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "account"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"getAccount"}

// AccountInfo is the result type of the account service getAccount method.
type AccountInfo struct {
	UserID     string
	Activated  bool
	Membership int
	Roles      []string
}
