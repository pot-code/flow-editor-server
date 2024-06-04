package design

//lint:ignore ST1001 dsl
import . "goa.design/goa/v3/dsl"

var FlowListItem = Type("FlowListItem", func() {
	Attribute("id", Int)
	Attribute("title", String)
	Attribute("created_at", String)
	Required("id", "title", "created_at")
})

var FlowDetail = Type("FlowDetail", func() {
	Attribute("id", Int)
	Attribute("title", String)
	Attribute("nodes", String)
	Attribute("edges", String)
	Attribute("created_at", String)
	Required("id", "title", "created_at")
})

var CreateFlowData = Type("CreateFlowData", func() {
	Attribute("title", String)
	Attribute("nodes", String)
	Attribute("edges", String)
	Required("title")
})

var UpdateFlowData = Type("UpdateFlowData", func() {
	Attribute("title", String)
	Attribute("nodes", String)
	Attribute("edges", String)
})

var _ = Service("flow", func() {
	Description("Flow 服务")
	HTTP(func() {
		Path("/flow")
	})
	Method("getFlowList", func() {
		Description("列出当前用户拥有的 flow")
		HTTP(func() {
			GET("/")
		})
		Result(ArrayOf(FlowListItem))
	})

	Method("getFlow", func() {
		Description("根据 flow id 获取 flow 详情")
		Payload(Int)
		HTTP(func() {
			GET("/{id}")
		})
		Result(FlowDetail)
	})

	Method("createFlow", func() {
		Description("创建 flow")
		Payload(CreateFlowData)
		HTTP(func() {
			POST("/")
			Response(StatusCreated)
		})
		Result(FlowDetail)
	})

	Method("updateFlow", func() {
		Description("更新 flow")
		Payload(func() {
			Attribute("data", UpdateFlowData)
			Attribute("id", Int)
		})
		HTTP(func() {
			PUT("/{id}")
			Body("data")
		})
		Result(FlowDetail)
	})

	Method("deleteFlow", func() {
		Description("删除 flow")
		Payload(Int)
		HTTP(func() {
			DELETE("/{id}")
			Response(StatusNoContent)
		})
	})
})
