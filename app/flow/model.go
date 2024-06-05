package flow

import (
	"gorm.io/gorm"
)

type Flow struct {
	gorm.Model
	Title string
	Nodes *string `gorm:"default:null"`
	Edges *string `gorm:"default:null"`
	Owner string
}

func (g *Flow) TableName() string {
	return "flows"
}
