package httperrors

import (
	"errors"
	dUser "ping-health/internal/domain/user"
	dMonitor "ping-health/internal/domain/monitor"
	"ping-health/internal/httpx"

)

func MapErrorsUser(err error) *httpx.AppError {
	switch{
		case errors.Is(err, dUser.ErrUserNotFound):
			return httpx.NotFound(err)
		case errors.Is(err, dUser.ErrExistingUser):
			return httpx.Conflict(err)
		case errors.Is(err, dUser.ErrHashedPassword):
			return httpx.Internal(err)
		case errors.Is(err, dUser.ErrInvalidCredentials):
			return httpx.Unauthorized(err)
		case errors.Is(err, dUser.ErrPasswordDontMatch):
			return httpx.BadRequest(err)
		default:
			return httpx.Internal(errors.New("erro interno -> user"))
	}
}

func MapErrorsMonitor(err error) *httpx.AppError {
	switch {
	case errors.Is(err, dMonitor.ErrInvalidData):
		return httpx.BadRequest(err)
	default:
		return httpx.Internal(errors.New("erro interno -> monitor"))
	}
}