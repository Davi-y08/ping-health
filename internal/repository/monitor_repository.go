package repository

import (
	"context"
	"errors"
	"ping-health/internal/domain/monitor"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MonitorRepository struct {
	db *gorm.DB
}

func NewMonitorRepository(db *gorm.DB) *MonitorRepository {
	return &MonitorRepository{db: db}
}

func (r *MonitorRepository) CreateMonitor(ctx context.Context, m monitor.Monitor) (error) {
	return r.db.WithContext(ctx).Model(&monitor.Monitor{}).Create(m).Error
}

func (r *MonitorRepository) GetMonitorById(ctx context.Context, id uuid.UUID) (*monitor.Monitor, error){
	var m *monitor.Monitor

	if err := r.db.WithContext(ctx).Model(&monitor.Monitor{}).First(&m, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, err
	}

	return m, nil
}

func (r *MonitorRepository) GetMonitorsByUser(ctx context.Context, user_id uuid.UUID) ([]monitor.Monitor, error) {
	var ms []monitor.Monitor

	if err := r.db.WithContext(ctx).Model(&monitor.Monitor{}).Where("user_id = ?", user_id).Find(&ms).Error; err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, err
	}

	return ms, nil
}