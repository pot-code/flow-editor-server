//go:build !goverter

package flow

import (
	"github.com/labstack/echo/v4"
	"github.com/zitadel/zitadel-go/v3/pkg/authorization"
)

type Controller struct {
	s *Service
	c Converter
}

func NewController(s *Service) *Controller {
	return &Controller{
		s: s,
		c: &ConverterImpl{},
	}
}

// deleteFlow
//
//	@Tags		flow
//	@Summary	delete flow
//	@Accept		json
//	@Param		id	path		string	true	"flow id"
//	@Success	204	{object}	nil
//	@Router		/flow/{id} [delete]
func (c *Controller) deleteFlow(ctx echo.Context) error {
	id := ctx.Param("id")
	o := authorization.Context[authorization.Ctx](ctx.Request().Context())
	if err := c.s.DeleteFlow(id, o.UserID()); err != nil {
		return err
	}
	return ctx.NoContent(204)
}

// getFlowList
//
//	@Tags		flow
//	@Summary	get flow list
//	@Produce	json
//	@Success	200	{array}	FlowListObjectOutput
//	@Router		/flow [get]
func (c *Controller) getFlowList(ctx echo.Context) error {
	o := authorization.Context[authorization.Ctx](ctx.Request().Context())
	flows, err := c.s.ListFlows(o.UserID())
	if err != nil {
		return err
	}
	return ctx.JSON(200, c.c.ConvertFlowModels(flows))
}

// getFlowDetail
//
//	@Tags		flow
//	@Summary	get flow detail
//	@Produce	json
//	@Param		id	path		string	true	"flow id"
//	@Success	200	{object}	FlowDetailOutput
//	@Router		/flow/{id} [get]
func (c *Controller) getFlowDetail(ctx echo.Context) error {
	id := ctx.Param("id")
	o := authorization.Context[authorization.Ctx](ctx.Request().Context())
	flow, err := c.s.GetFlow(id, o.UserID())
	if err != nil {
		return err
	}
	return ctx.JSON(200, c.c.ConvertFlowModel(*flow))
}

// createFlow
//
//	@Tags		flow
//	@Summary	create flow
//	@Accept		json
//	@Produce	json
//	@Param		payload	body		CreateFlowInput	true	"payload"
//	@Success	201		{object}	FlowDetailOutput
//	@Router		/flow [post]
func (c *Controller) createFlow(ctx echo.Context) error {
	var payload CreateFlowInput
	if err := ctx.Bind(&payload); err != nil {
		return err
	}

	o := authorization.Context[authorization.Ctx](ctx.Request().Context())
	payload.Owner = o.UserID()
	m, err := c.s.CreateFlow(&payload)
	if err != nil {
		return err
	}
	return ctx.JSON(201, c.c.ConvertFlowModel(*m))
}

// updateFlow
//
//	@Tags		flow
//	@Summary	update flow
//	@Accept		json
//	@Produce	json
//	@Param		id		path		string			true	"flow id"
//	@Param		payload	body		UpdateFlowInput	true	"payload"
//	@Success	200		{object}	FlowDetailOutput
//	@Router		/flow/{id} [put]
func (c *Controller) updateFlow(ctx echo.Context) error {
	id := ctx.Param("id")

	var payload UpdateFlowInput
	if err := ctx.Bind(&payload); err != nil {
		return err
	}

	o := authorization.Context[authorization.Ctx](ctx.Request().Context())
	payload.Owner = o.UserID()
	m, err := c.s.UpdateFlow(id, &payload)
	if err != nil {
		return err
	}
	return ctx.JSON(200, c.c.ConvertFlowModel(*m))
}

func (c *Controller) RegisterHandlers(e *echo.Echo) {
	r := e.Group("/flow")
	r.GET("", c.getFlowList)
	r.GET("/:id", c.getFlowDetail)
	r.POST("", c.createFlow)
	r.PUT("/:id", c.updateFlow)
	r.DELETE("/:id", c.deleteFlow)
}
