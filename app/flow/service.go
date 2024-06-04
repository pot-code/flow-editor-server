package flow

import (
	"context"
	"flow-editor-server/gen/flow"

	"github.com/open-policy-agent/opa/sdk"
	"github.com/zitadel/zitadel-go/v3/pkg/authorization"
	"gorm.io/gorm"
)

type Service struct {
	db  *gorm.DB
	c   Converter
	opa *sdk.OPA
}

// CopyFlow implements flow.Service.
func (s *Service) CopyFlow(ctx context.Context, copyId string) (err error) {
	auth := authorization.Context[authorization.Ctx](ctx)
	var m FlowModel
	if err := s.db.First(&m, copyId).Error; err != nil {
		return err
	}

	m.Owner = auth.UserID()
	return s.db.Model(&FlowModel{}).Omit("id").Create(&m).Error
}

// CreateFlow implements flow.Service.
func (s *Service) CreateFlow(ctx context.Context, data *flow.CreateFlowData) (res *flow.FlowDetail, err error) {
	auth := authorization.Context[authorization.Ctx](ctx)
	m := &FlowModel{
		Title: *data.Title,
		Nodes: data.Nodes,
		Edges: data.Edges,
		Owner: auth.UserID(),
	}
	if err := s.db.Create(m).Error; err != nil {
		return nil, err
	}
	return s.c.FlowModelToFlowDetail(m), nil
}

// DeleteFlow implements flow.Service.
func (s *Service) DeleteFlow(ctx context.Context, id string) (err error) {
	auth := authorization.Context[authorization.Ctx](ctx)
	return s.db.Where("id = ? AND owner = ?", id, auth.UserID()).Delete(&FlowModel{}).Error
}

// GetFlow implements flow.Service.
func (s *Service) GetFlow(ctx context.Context, id string) (res *flow.FlowDetail, err error) {
	auth := authorization.Context[authorization.Ctx](ctx)
	var flow FlowModel
	if err := s.db.Where("id = ? AND owner = ?", id, auth.UserID()).First(&flow).Error; err != nil {
		return nil, err
	}
	return s.c.FlowModelToFlowDetail(&flow), nil
}

// GetFlowList implements flow.Service.
func (s *Service) GetFlowList(ctx context.Context) (res []*flow.FlowListItem, err error) {
	auth := authorization.Context[authorization.Ctx](ctx)
	var flows []*FlowModel
	if err := s.db.Find(&flows, "owner = ?", auth.UserID()).Error; err != nil {
		return nil, err
	}
	return s.c.FlowModelsToFlowList(flows), nil
}

// UpdateFlow implements flow.Service.
func (s *Service) UpdateFlow(ctx context.Context, payload *flow.UpdateFlowPayload) (res *flow.FlowDetail, err error) {
	auth := authorization.Context[authorization.Ctx](ctx)
	data := payload.Data

	var model FlowModel
	if err := s.db.First(&model, "id = ? AND owner = ?", payload.ID, auth.UserID()).Error; err != nil {
		return nil, err
	}
	if data.Edges != nil {
		model.Edges = data.Edges
	}
	if data.Nodes != nil {
		model.Nodes = data.Nodes
	}
	if data.Title != nil {
		model.Title = *data.Title
	}
	if err := s.db.Save(&model).Error; err != nil {
		return nil, err
	}
	return s.c.FlowModelToFlowDetail(&model), nil
}

var _ flow.Service = (*Service)(nil)

func NewService(db *gorm.DB, c Converter, opa *sdk.OPA) *Service {
	return &Service{db: db, c: c, opa: opa}
}
