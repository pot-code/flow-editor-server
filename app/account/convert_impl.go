// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.
//go:build !goverter

package account

type ConverterImpl struct{}

func (c *ConverterImpl) ConvertAccountModel(source AccountModel) AccountObject {
	var accountAccountObject AccountObject
	accountAccountObject.Activated = source.Activated
	accountAccountObject.Membership = int(source.Membership)
	return accountAccountObject
}
