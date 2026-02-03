package monitor

import (
	"ping-health/internal/domain/user"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Monitor struct {
	gorm.Model
	URL 		string 			`json:"url"`
	Interval	time.Duration	`json:"interval"`
	UserID		uuid.UUID		`json:"user_id"`
	User 		user.User		`gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}