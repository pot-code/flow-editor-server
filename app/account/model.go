package account

import "gorm.io/gorm"

type MembershipType int

const (
	MembershipTypeFree MembershipType = iota
	MembershipTypePro
	MembershipTypeEnterprise
)

type AccountModel struct {
	gorm.Model
	UserId     string `gorm:"unique;not null"`
	Membership MembershipType
	Activated  bool
}

func (*AccountModel) TableName() string {
	return "accounts"
}
