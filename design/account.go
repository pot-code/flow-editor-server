package design

//lint:ignore ST1001 dsl
import . "goa.design/goa/v3/dsl"

var AccountInfo = Type("AccountInfo", func() {
	Attribute("activated", Boolean)
	Attribute("membership", Int)
	Required("activated", "membership")
})

var _ = Service("account", func() {
	Description("Account service")
	HTTP(func() {
		Path("/account")
	})
	Method("getAccount", func() {
		Description("Get account")
		Result(AccountInfo)
		HTTP(func() {
			GET("/")
		})
	})
})
