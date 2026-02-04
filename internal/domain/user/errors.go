package user

import "errors"

var (
	ErrUserNotFound = errors.New("usuário não encontrado")
	ErrUserInvalidData = errors.New("dados inválidos")
	ErrHashedPassword = errors.New("erro ao hashear senha")
	ErrExistingUser = errors.New("usuário já existente")
)