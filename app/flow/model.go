package flow

import (
	"flow-editor-server/internal/authz"
	"strconv"

	"gorm.io/gorm"
)

type Flow struct {
	gorm.Model
	Title string
	Nodes *string `gorm:"default:null"`
	Edges *string `gorm:"default:null"`
	Owner string
}

func (m *Flow) Kind() string {
	return "flow"
}

func (m *Flow) OwnerID() string {
	return m.Owner
}

func (m *Flow) ResourceID() string {
	return strconv.Itoa(int(m.ID))
}

var _ authz.Resource = (*Flow)(nil)
