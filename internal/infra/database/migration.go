package database

import (
	"ping-health/internal/domain/user"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&user.User{})
}