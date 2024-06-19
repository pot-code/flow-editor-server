package flow

import (
	"gorm.io/gorm"
)

type Flow struct {
	gorm.Model
	Title string
	Data  *string `gorm:"default:null"`
	Owner string
}
