package flow

import "github.com/labstack/echo/v4"

type Controller struct {
	s *Service
	c *ConverterImpl
}

func NewController(s *Service) *Controller {
	return &Controller{
		s: s,
		c: &ConverterImpl{},
	}
}

// DeleteFlowId implements ServerInterface.
func (c *Controller) DeleteFlowId(ctx echo.Context, id string) error {
	if err := c.s.DeleteFlow(id); err != nil {
		return err
	}
	return ctx.NoContent(204)
}

// GetFlow implements ServerInterface.
func (c *Controller) GetFlow(ctx echo.Context) error {
	flows, err := c.s.ListFlows()
	if err != nil {
		return err
	}
	return ctx.JSON(200, c.c.ConvertListFlow(flows))
}

// GetFlowId implements ServerInterface.
func (c *Controller) GetFlowId(ctx echo.Context, id string) error {
	flow, err := c.s.GetFlow(id)
	if err != nil {
		return err
	}
	return ctx.JSON(200, c.c.ConvertFlowDetail(*flow))
}

// PostFlow implements ServerInterface.
func (c *Controller) PostFlow(ctx echo.Context) error {
	panic("unimplemented")
}

// PutFlow implements ServerInterface.
func (c *Controller) PutFlow(ctx echo.Context) error {
	panic("unimplemented")
}

var _ ServerInterface = (*Controller)(nil)
