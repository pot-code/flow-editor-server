//go:build !goverter

package account

import (
	"context"
	"errors"
	"flow-editor-server/gen/account"

	"github.com/zitadel/zitadel-go/v3/pkg/authorization"
	"gorm.io/gorm"
)

type service struct {
	db *gorm.DB
	c  Converter
}

// GetAccount implements account.Service.
func (s *service) GetAccount(ctx context.Context) (res *account.AccountInfo, err error) {
	auth := authorization.Context[authorization.Ctx](ctx)

	var a AccountModel
	err = s.db.First(&a, &AccountModel{UserId: auth.UserID()}).Error
	if err == nil {
		res = s.c.ConvertAccountModel(a)
		return
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}

	a.UserId = auth.UserID()
	a.Membership = MembershipTypeFree
	a.Activated = true
	if err = s.db.Create(&a).Error; err != nil {
		return
	}
	return s.c.ConvertAccountModel(a), nil
}

var _ account.Service = (*service)(nil)

func NewService(db *gorm.DB, c Converter) *service {
	return &service{db: db, c: c}
}
