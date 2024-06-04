package account

import "flow-editor-server/gen/account"

// goverter:converter
// goverter:output:file convert_impl.go
// goverter:output:package flow-editor-server/app/account
type Converter interface {
	FromAccountModelToAccountInfo(m AccountModel) *account.AccountInfo
}
