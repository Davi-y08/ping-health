package ping

import (
	"ping-health/internal/domain/monitor"
	"time"

	"github.com/google/uuid"
)

type Ping struct {
	ID        	uuid.UUID 		`gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	StatusCode 	int 			`json:"status_code"`
	Status	    bool			`json:"status"`
	URL 		string 			`json:"url"`
	MonitorID	uuid.UUID		`gorm:"type:uuid;not null"`
	Montitor 	monitor.Monitor	`gorm:"foreignKey:MonitorID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}