package monitor

import (
	dMonitor "ping-health/internal/domain/monitor"
	"github.com/google/uuid"
)

type CreateMonitorDto struct {
	URL      	string 		`json:"url"`
	Interval 	int    		`json:"interval"`
	UserID   	uuid.UUID 	`json:"user_id"`
}

func ValidateDto(dto CreateMonitorDto) (*dMonitor.Monitor, error){
	if dto.URL == "" || dto.Interval == 0 || dto.UserID.String() == ""{
		return nil, dMonitor.ErrInvalidData
	}

	new_monitor := &dMonitor.Monitor{
		URL: dto.URL,
		Interval: dto.Interval,
		UserID: dto.UserID,
	}

	return new_monitor, nil
}