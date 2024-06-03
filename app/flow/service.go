package flow

import (
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s *Service) CreateFlow(flow *CreateFlowInput) (*FlowModel, error) {
	if err := flow.Validate(); err != nil {
		return nil, err
	}

	m := &FlowModel{
		Title: flow.Title,
		Nodes: flow.Nodes,
		Edges: flow.Edges,
		Owner: flow.Owner,
	}
	if err := s.db.Create(m).Error; err != nil {
		return nil, err
	}
	return m, nil
}

func (s *Service) GetFlow(id string, owner string) (*FlowModel, error) {
	var flow FlowModel
	if err := s.db.Where("id = ? AND owner = ?", id, owner).First(&flow).Error; err != nil {
		return nil, err
	}
	return &flow, nil
}

func (s *Service) DeleteFlow(id string, owner string) error {
	return s.db.Where("id = ? AND owner = ?", id, owner).Delete(&FlowModel{}).Error
}

func (s *Service) UpdateFlow(id string, flow *UpdateFlowInput) (*FlowModel, error) {
	if err := flow.Validate(); err != nil {
		return nil, err
	}

	var model FlowModel
	if err := s.db.First(&model, "id = ? AND owner = ?", id, flow.Owner).Error; err != nil {
		return nil, err
	}
	if err := s.db.Model(&model).Omit("id").Updates(&flow).Error; err != nil {
		return nil, err
	}
	return &model, nil
}

func (s *Service) ListFlows(owner string) ([]FlowModel, error) {
	var flows []FlowModel
	if err := s.db.Find(&flows, "owner = ?", owner).Error; err != nil {
		return nil, err
	}
	return flows, nil
}
