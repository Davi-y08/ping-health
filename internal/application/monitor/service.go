package monitor

import (
	"context"
	repo "ping-health/internal/repository"
	shared "ping-health/internal/application"
)

type MonitorService struct{
	repo *repo.MonitorRepository
}

func NewMonitorService(repo *repo.MonitorRepository) *MonitorService {
	return &MonitorService{repo: repo}
}

func (s *MonitorService) CreateMonitorService(ctx context.Context, dto CreateMonitorDto) (error) {
	new_monitor, err := ValidateDto(dto)

	if err != nil {
		return err	
	}

	if err := s.repo.CreateMonitor(ctx, new_monitor); err != nil{
		return shared.ErrInDataBase
	}

	return nil
}