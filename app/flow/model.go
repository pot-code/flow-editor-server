package flow

import (
	"gorm.io/gorm"
)

type FlowModel struct {
	gorm.Model
	Title string
	Nodes *string
	Edges *string
}

func (g *FlowModel) TableName() string {
	return "flows"
}
