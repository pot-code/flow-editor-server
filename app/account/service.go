//go:build !goverter

package account

import (
	"context"
	"errors"
	"flow-editor-server/gen/account"
	"flow-editor-server/internal/authn"

	"gorm.io/gorm"
)

type service struct {
	db *gorm.DB
	c  Converter
}

// GetAccount implements account.Service.
func (s *service) GetAccount(ctx context.Context) (*account.AccountInfo, error) {
	token := authn.FromContext(ctx)

	var a Account
	err := s.db.Preload("Roles").First(&a, &Account{UserID: token.Subject()}).Error
	if err == nil {
		return s.c.AccountToAccountInfo(a), nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return s.createAccount(ctx)
}

func (s *service) createAccount(ctx context.Context) (*account.AccountInfo, error) {
	token := authn.FromContext(ctx)

	var role Role
	if err := s.db.Where("name = ?", "user").First(&role).Error; err != nil {
		return nil, err
	}

	var a Account
	a.UserID = token.Subject()
	a.Membership = MembershipTypeFree
	a.Activated = true
	if err := s.db.Create(&a).Error; err != nil {
		return nil, err
	}
	if err := s.db.Omit("Roles.*").Model(&a).Association("Roles").Append(&role); err != nil {
		return nil, err
	}
	return s.c.AccountToAccountInfo(a), nil
}

var _ account.Service = (*service)(nil)

func NewService(db *gorm.DB, c Converter) *service {
	return &service{db: db, c: c}
}
