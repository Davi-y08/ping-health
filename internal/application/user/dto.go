package user

import (
	"net/mail"
	"ping-health/internal/domain/user"
	security "ping-health/internal/infra/security"
)

type CreateUserDto struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	PassWord        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func validateEmail(str string) bool {
    _, err := mail.ParseAddress(str)
    return err == nil
}

func ValidateDto(dto CreateUserDto) (*user.User, error) {
	if !validateEmail(dto.Email) || len(dto.Name) < 5 || len(dto.PassWord) < 6{
		return nil, user.ErrUserInvalidData
	}

	hash, err := security.HashPassword(dto.PassWord)

	if err != nil {
		return nil, user.ErrHashedPassword
	}

	new_user := &user.User{
		Email: dto.Email,
		PasswordHash: hash,
		Name: dto.Name,
		Role: user.DefaultRole,
	}

	return new_user, nil
}