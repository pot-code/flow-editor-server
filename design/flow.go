package design

//lint:ignore ST1001 dsl
import . "goa.design/goa/v3/dsl"

var FlowListItemData = Type("FlowListItemData", func() {
	Attribute("id", Int, "flow id")
	Attribute("title", String, "flow 标题")
	Attribute("created_at", String, "flow 创建时间")
	Required("id", "title", "created_at")
})

var FlowDetailData = Type("FlowDetailData", func() {
	Attribute("id", Int, "flow id")
	Attribute("title", String, "flow 标题")
	Attribute("nodes", String, "flow 节点")
	Attribute("edges", String, "flow 边")
	Attribute("created_at", String, "flow 创建时间")
	Required("id", "title", "created_at")
})

var CreateFlowData = Type("CreateFlowData", func() {
	Attribute("title", String, "flow 标题", func() {
		Meta("struct:tag:json", "title")
		Meta("struct:tag:validate", "required,min=1,max=32")
	})
	Attribute("nodes", String, "flow 节点")
	Attribute("edges", String, "flow 边")
})

var UpdateFlowData = Type("UpdateFlowData", func() {
	Attribute("title", String, "flow 标题", func() {
		Meta("struct:tag:json", "title")
		Meta("struct:tag:validate", "required,min=1,max=32")
	})
	Attribute("nodes", String, "flow 节点")
	Attribute("edges", String, "flow 边")
})

var QueryFlowListParams = Type("QueryFlowListParams", func() {
	Attribute("name", String, "流程名")
})

var _ = Service("flow", func() {
	Description("Flow 服务")
	HTTP(func() {
		Path("/flow")
	})

	Method("getFlowList", func() {
		Description("列出当前用户拥有的 flow")
		Payload(QueryFlowListParams)
		HTTP(func() {
			GET("/")
			Param("name")
		})
		Result(ArrayOf(FlowListItemData))
	})

	Method("getFlow", func() {
		Description("根据 flow id 获取 flow 详情")
		Payload(String, "要获取的 flow id")
		HTTP(func() {
			GET("/{id}")
		})
		Result(FlowDetailData)
	})

	Method("createFlow", func() {
		Description("创建 flow")
		Payload(CreateFlowData)
		HTTP(func() {
			POST("/")
			Response(StatusCreated)
		})
		Result(FlowDetailData)
	})

	Method("updateFlow", func() {
		Description("更新 flow")
		Payload(func() {
			Attribute("data", UpdateFlowData)
			Attribute("id", String, "要更新的 flow id")
		})
		HTTP(func() {
			PUT("/{id}")
			Body("data")
		})
		Result(FlowDetailData)
	})

	Method("deleteFlow", func() {
		Description("删除 flow")
		Payload(String, "要删除的 flow id")
		HTTP(func() {
			DELETE("/{id}")
			Response(StatusNoContent)
		})
	})

	Method("copyFlow", func() {
		Description("复制 flow")
		Payload(String, "要复制的 flow id")
		HTTP(func() {
			POST("/{id}/copy")
			Response(StatusCreated)
		})
	})
})
