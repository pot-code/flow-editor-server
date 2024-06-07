package account

import (
	"flow-editor-server/gen/account"
)

// goverter:converter
// goverter:output:file convert_impl.go
// goverter:output:package flow-editor-server/app/account
// goverter:extend RoleToString
type Converter interface {
	AccountToAccountInfo(m Account) *account.AccountInfo
}

func RoleToString(role Role) string {
	return role.Name
}
