package flow

import "gorm.io/gorm"

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s *Service) CreateFlow(flow *CreateFlow) error {
	return s.db.Create(flow).Error
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
