package user

import (
	"context"
	"errors"
	shared "ping-health/internal/application"
	"ping-health/internal/domain/user"
	security "ping-health/internal/infra/security"
	repo "ping-health/internal/repository"

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

	if err := s.repo.CreateUser(ctx, u); err != nil{
		return shared.ErrInDataBase
	}

	return nil
}

func (s *UserService) Login(ctx context.Context, dto LoginDto) (*user.User, error){
	if dto.Email == "" || dto.Password == ""{
		return nil, user.ErrInvalidCredentials
	}

	u, err := s.repo.GetUserByEmail(ctx, dto.Email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return nil, user.ErrInvalidCredentials
		}

		return nil, shared.ErrInDataBase
	}

	if !security.CheckPassword(u.PasswordHash, dto.Password){
		return nil, user.ErrInvalidCredentials
	}

	return u, nil
}