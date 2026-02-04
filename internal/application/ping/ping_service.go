package ping

import repo "ping-health/internal/repository"

type PingService struct{
	repo *repo.PingRepository
}

func NewPingService(repo *repo.PingRepository) *PingService {
	return &PingService{repo: repo}
}