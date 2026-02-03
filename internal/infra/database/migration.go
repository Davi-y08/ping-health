package database

import (
	"ping-health/internal/domain/monitor"
	"ping-health/internal/domain/ping"
	"ping-health/internal/domain/user"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&monitor.Monitor{})
	db.AutoMigrate(&ping.Ping{})
}