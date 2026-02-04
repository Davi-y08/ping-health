package user

import (
	"context"
	"ping-health/internal/domain/user"
	repo "ping-health/internal/repository"
	shared "ping-health/internal/application"
	"gorm.io/gorm"
)

type UserService struct{
	repo *repo.UserRepository
}

func NewUserService(repo *repo.UserRepository) *UserService{
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, dto CreateUserDto) (error){
	u, err := ValidateDto(dto)

	if err != nil {
		return err
	}

	existing, err := s.repo.GetUserByEmail(ctx, u.Email)

	if err != gorm.ErrRecordNotFound || existing != nil{
		return user.ErrExistingUser
	}

	if err := s.repo.CreateUser(ctx, *u); err != nil{
		return shared.ErrInDataBase
	}

	return nil
}