package design

//lint:ignore ST1001 dsl
import . "goa.design/goa/v3/dsl"

var AccountInfo = Type("AccountInfo", func() {
	Attribute("user_id", String)
	Attribute("activated", Boolean)
	Attribute("membership", Int)
	Attribute("roles", ArrayOf(String))
	Required("user_id", "activated", "membership")
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
