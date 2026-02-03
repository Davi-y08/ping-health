package repository

import (
	"context"
	"errors"
	"ping-health/internal/domain/ping"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PingRepository struct {
	db *gorm.DB
}

func NewPingRepository(db *gorm.DB) *PingRepository {
	return &PingRepository{db: db}
}

func (r *PingRepository) CreatePing(ctx context.Context, p ping.Ping) (error) {
	return r.db.WithContext(ctx).Model(&ping.Ping{}).Create(p).Error
}

func (r *PingRepository) GetPingById(ctx context.Context, id uuid.UUID) (*ping.Ping, error){
	var p *ping.Ping

	if err := r.db.WithContext(ctx).Model(&ping.Ping{}).First(p, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, err
	}

	return p, nil
}

func (r *PingRepository) GetPingsByMonitor(ctx context.Context, monitor_id uuid.UUID) ([]ping.Ping, error) {
	var ps []ping.Ping

	if err := r.db.WithContext(ctx).Model(&ping.Ping{}).Where("monitor_id = ?", monitor_id).Find(ps).Error; err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			return nil, gorm.ErrRecordNotFound
		}

		return nil, err
	}

	return ps, nil
}

func (r *PingRepository) DeletePing(ctx context.Context, id uuid.UUID) (error) {
	result := r.db.WithContext(ctx).Model(&ping.Ping{}).Unscoped().Delete(&ping.Ping{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}