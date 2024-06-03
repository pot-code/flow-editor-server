package design

import (
	//lint:ignore ST1001 dsl
	. "goa.design/goa/v3/dsl"
)

var AccountOutput = Type("AccountOutput", func() {
	Attribute("activated", Boolean)
	Attribute("membership", Int)
})

var _ = Service("account", func() {
	Description("Account service")
	HTTP(func() {
		Path("/account")
	})
	Method("getAccount", func() {
		Description("Get account")
		Result(AccountOutput)
		HTTP(func() {
			GET("/")
		})
	})
})
