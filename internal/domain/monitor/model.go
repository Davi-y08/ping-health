package monitor

import (
	"ping-health/internal/domain/user"
	"time"
	"github.com/google/uuid"
)

type Monitor struct {
	ID        	uuid.UUID 		`gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	URL 		string 			`json:"url"`
	Interval	time.Duration	`json:"interval"`
	UserID		uuid.UUID		`gorm:"type:uuid;not null"`
	User 		user.User		`gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}