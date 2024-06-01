package account

import "gorm.io/gorm"

type AccountModel struct {
	gorm.Model
	UserId     string
	Membership int
	Activated  bool
}
