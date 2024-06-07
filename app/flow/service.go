package flow

import (
	"context"
	"flow-editor-server/gen/flow"
	"flow-editor-server/internal/authz"

	"github.com/cerbos/cerbos-sdk-go/cerbos"
	"github.com/zitadel/zitadel-go/v3/pkg/authorization"
	"gorm.io/gorm"
)

type service struct {
	db *gorm.DB
	c  Converter
	a  *authz.Authz[*Flow]
}

// CopyFlow implements flow.Service.
func (s *service) CopyFlow(ctx context.Context, copyId string) (err error) {
	auth := authorization.Context[authorization.Ctx](ctx)
	var m Flow
	if err := s.db.First(&m, copyId).Error; err != nil {
		return err
	}

	m.Owner = auth.UserID()
	return s.db.Model(&Flow{}).Omit("id").Create(&m).Error
}

// CreateFlow implements flow.Service.
func (s *service) CreateFlow(ctx context.Context, data *flow.CreateFlowData) (res *flow.FlowDetailData, err error) {
	a := authz.Context(ctx)
	m := &Flow{
		Title: *data.Title,
		Nodes: data.Nodes,
		Edges: data.Edges,
		Owner: a.UserID,
	}
	if err := s.db.Create(m).Error; err != nil {
		return nil, err
	}
	return s.c.FlowModelToFlowDetail(m), nil
}

// DeleteFlow implements flow.Service.
func (s *service) DeleteFlow(ctx context.Context, id string) (err error) {
	var f Flow
	if err := s.db.First(&f, id).Error; err != nil {
		return err
	}
	if err := s.a.CheckPermission(ctx, &f, "delete"); err != nil {
		return err
	}
	return s.db.Delete(&f).Error
}

// GetFlow implements flow.Service.
func (s *service) GetFlow(ctx context.Context, id string) (res *flow.FlowDetailData, err error) {
	a := authz.Context(ctx)
	var flow Flow
	if err := s.db.Where("id = ? AND owner = ?", id, a.UserID).First(&flow).Error; err != nil {
		return nil, err
	}
	return s.c.FlowModelToFlowDetail(&flow), nil
}

// GetFlowList implements flow.Service.
func (s *service) GetFlowList(ctx context.Context) (res []*flow.FlowListItemData, err error) {
	a := authz.Context(ctx)
	var flows []*Flow
	if err := s.db.Find(&flows, "owner = ?", a.UserID).Error; err != nil {
		return nil, err
	}
	return s.c.FlowModelsToFlowList(flows), nil
}

// UpdateFlow implements flow.Service.
func (s *service) UpdateFlow(ctx context.Context, payload *flow.UpdateFlowPayload) (res *flow.FlowDetailData, err error) {
	data := payload.Data

	var model Flow
	if err := s.db.First(&model, "id = ?", payload.ID).Error; err != nil {
		return nil, err
	}
	if err := s.a.CheckPermission(ctx, &model, "update"); err != nil {
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

var _ flow.Service = (*service)(nil)

func NewService(db *gorm.DB, c Converter, cb *cerbos.GRPCClient) *service {
	return &service{db: db, c: c, a: authz.NewAuthz[*Flow](cb)}
}
