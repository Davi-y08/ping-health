package repository

import (
	"context"
	"errors"
	"os/user"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository{
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, u user.User) (error){
	return r.db.WithContext(ctx).Model(&user.User{}).Create(u).Error
}

func (r *UserRepository) GetUserById(ctx context.Context, id uuid.UUID) (*user.User, error){
	var u *user.User
	
	if err := r.db.WithContext(ctx).Model(&user.User{}).First(u, id).Error; err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, err
	}

	return u, nil
}