package ping

import (
	"ping-health/internal/domain/monitor"
	"gorm.io/gorm"
)

type Ping struct {
	gorm.Model
	StatusCode 	int 			`json:"status_code"`
	Status	    bool			`json:"status"`
	URL 		string 			`json:"url"`
	MonitorID	uint			`json:"monitor_id"`
	Montitor 	monitor.Monitor	`gorm:"foreignKey:MonitorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}