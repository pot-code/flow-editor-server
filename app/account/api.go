//go:build !goverter

package account

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/zitadel/zitadel-go/v3/pkg/authorization"
	"gorm.io/gorm"
)

type Controller struct {
	db *gorm.DB
	c  Converter
}

// getAccount
//
//	@Tags		account
//	@Summary	get account
//	@Produce	json
//	@Success	200	{object}	AccountOutput
//	@Router		/account [get]
func (c *Controller) getAccount(ctx echo.Context) error {
	auth := authorization.Context[authorization.Ctx](ctx.Request().Context())

	var a AccountModel
	err := c.db.First(&a, &AccountModel{UserId: auth.UserID()}).Error
	if err == nil {
		return ctx.JSON(200, c.c.ConvertAccountModel(a))
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	a.UserId = auth.UserID()
	a.Membership = MembershipTypeFree
	a.Activated = true
	if err := c.db.Create(&a).Error; err != nil {
		return err
	}
	return ctx.JSON(200, c.c.ConvertAccountModel(a))
}

func (c *Controller) RegisterHandlers(e *echo.Echo) {
	r := e.Group("/account")
	r.GET("", c.getAccount)
}

func NewController(db *gorm.DB) *Controller {
	return &Controller{db: db, c: &ConverterImpl{}}
}
