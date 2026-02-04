package monitor

import repo "ping-health/internal/repository"

type MonitorService struct{
	repo *repo.MonitorRepository
}

func NewMonitorService(repo *repo.MonitorRepository) *MonitorService {
	return &MonitorService{repo: repo}
}

