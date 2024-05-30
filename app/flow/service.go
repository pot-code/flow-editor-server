package flow

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
)

type CreateFlowData struct {
	Title string  `json:"title"`
	Nodes *string `json:"nodes"`
	Edges *string `json:"edges"`
}

func (c *CreateFlowData) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.Title, validation.Required, validation.Length(1, 32)),
	)
}

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s *Service) CreateFlow(flow *CreateFlowData) error {
	if err := flow.Validate(); err != nil {
		return err
	}
	return s.db.Create(&FlowModel{
		Title: flow.Title,
		Nodes: flow.Nodes,
		Edges: flow.Edges,
	}).Error
}

func (s *Service) GetFlow(id string) (*FlowDetail, error) {
	var flow FlowDetail
	if err := s.db.Where("id = ?", id).First(&flow).Error; err != nil {
		return nil, err
	}
	return &flow, nil
}

func (s *Service) DeleteFlow(id string) error {
	return s.db.Where("id = ?", id).Delete(&FlowModel{}).Error
}

func (s *Service) UpdateFlow(flow *FlowModel) error {
	return s.db.Save(flow).Error
}

func (s *Service) ListFlows() ([]FlowListItem, error) {
	var flows []FlowListItem
	if err := s.db.Model(&FlowModel{}).Find(&flows).Error; err != nil {
		return nil, err
	}
	return flows, nil
}
