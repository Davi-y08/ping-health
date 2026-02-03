package monitor

import (
	"ping-health/internal/domain/user"
	"time"
	"gorm.io/gorm"
)

type Monitor struct {
	gorm.Model
	URL 		string 			`json:"url"`
	Interval	time.Duration	`json:"interval"`
	UserID		uint			`json:"user_id"`
	User 		user.User		`gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}