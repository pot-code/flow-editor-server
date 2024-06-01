package account

//go:generate go run github.com/jmattheis/goverter/cmd/goverter gen .

// goverter:converter
// goverter:output:file convert_impl.go
// goverter:output:package openapi-go-demo/app/account
type Converter interface {
	ConvertAccountModel(m AccountModel) AccountObject
}
