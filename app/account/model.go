package account

import "gorm.io/gorm"

type MembershipType int

const (
	MembershipTypeFree MembershipType = iota
	MembershipTypePro
	MembershipTypeEnterprise
)

type Account struct {
	gorm.Model
	UserId     string `gorm:"unique;not null"`
	Membership MembershipType
	Activated  bool
	Roles      []Role `gorm:"many2many:account_roles;"`
}

type Role struct {
	gorm.Model
	Name        string
	Description string
}
