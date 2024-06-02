package flow

import (
	"gorm.io/gorm"
)

type FlowModel struct {
	gorm.Model
	Title string
	Nodes *string `gorm:"default:null"`
	Edges *string `gorm:"default:null"`
	Owner string
}

func (g *FlowModel) TableName() string {
	return "flows"
}
