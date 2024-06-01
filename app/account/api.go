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

// GetAccount implements ServerInterface.
func (c *Controller) GetAccount(ctx echo.Context) error {
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

var _ ServerInterface = (*Controller)(nil)

func NewController(db *gorm.DB) *Controller {
	return &Controller{db: db, c: &ConverterImpl{}}
}
