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

// DeleteFlowId implements ServerInterface.
func (c *Controller) DeleteFlowId(ctx echo.Context, id string) error {
	o := authorization.Context[authorization.Ctx](ctx.Request().Context())
	if err := c.s.DeleteFlow(id, o.UserID()); err != nil {
		return err
	}
	return ctx.NoContent(204)
}

// GetFlow implements ServerInterface.
func (c *Controller) GetFlow(ctx echo.Context) error {
	o := authorization.Context[authorization.Ctx](ctx.Request().Context())
	flows, err := c.s.ListFlows(o.UserID())
	if err != nil {
		return err
	}
	return ctx.JSON(200, c.c.ConvertSliceFlowListItem(flows))
}

// GetFlowId implements ServerInterface.
func (c *Controller) GetFlowId(ctx echo.Context, id string) error {
	o := authorization.Context[authorization.Ctx](ctx.Request().Context())

	flow, err := c.s.GetFlow(id, o.UserID())
	if err != nil {
		return err
	}
	return ctx.JSON(200, c.c.ConvertFlowModel(*flow))
}

// PostFlow implements ServerInterface.
func (c *Controller) PostFlow(ctx echo.Context) error {
	var payload PostFlowJSONRequestBody
	if err := ctx.Bind(&payload); err != nil {
		return err
	}

	o := authorization.Context[authorization.Ctx](ctx.Request().Context())
	data := c.c.ConvertPostFlowJSONRequestBody(payload)
	data.Owner = o.UserID()
	m, err := c.s.CreateFlow(&data)
	if err != nil {
		return err
	}
	return ctx.JSON(201, c.c.ConvertFlowModel(*m))
}

// PutFlowId implements ServerInterface.
func (c *Controller) PutFlowId(ctx echo.Context, id int) error {
	var payload PutFlowIdJSONRequestBody
	if err := ctx.Bind(&payload); err != nil {
		return err
	}

	o := authorization.Context[authorization.Ctx](ctx.Request().Context())
	data := c.c.ConvertPutFlowJSONRequestBody(payload)
	data.Owner = o.UserID()
	m, err := c.s.UpdateFlow(id, &data)
	if err != nil {
		return err
	}
	return ctx.JSON(200, c.c.ConvertFlowModel(*m))
}

var _ ServerInterface = (*Controller)(nil)
