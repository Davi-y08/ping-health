package monitor

import (
	"context"
	"fmt"
	"net/http"
	shared "ping-health/internal/application"
	"ping-health/internal/domain/monitor"
	repo "ping-health/internal/repository"
	"time"
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

	go s.startMonitor(new_monitor)

	return nil
}

func (s *MonitorService) startMonitor(m *monitor.Monitor){
	ticker := time.NewTicker(time.Duration(m.Interval) * time.Second)
	defer ticker.Stop()

	for {
		s.check(m.URL)
		<-ticker.C
	}
}

func (s *MonitorService) check(url string){
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)

	if err != nil {
		fmt.Println("off")
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		fmt.Println(url, " online")
	}else{
		fmt.Println(url, "offline", resp.StatusCode)
	}
}