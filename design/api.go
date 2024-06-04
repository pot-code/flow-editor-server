package design

//lint:ignore ST1001 dsl
import . "goa.design/goa/v3/dsl"

var _ = API("flow-editor-server", func() {
	Title("流程编辑器 API")
})
