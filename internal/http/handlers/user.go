package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	appUser "ping-health/internal/application/user"
	service "ping-health/internal/application/user"
	"ping-health/internal/httpx"
	mapErrors "ping-health/internal/http/http_errors"
)

type UserHandler struct{
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service }
}

func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		return httpx.MethodNotAllowed(errors.New("method not allowed"))
	}

	var dto appUser.CreateUserDto

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil{
		return httpx.BadRequest(errors.New("corpo inválido"))
	}

	if create_err := h.service.CreateUser(r.Context(), dto); create_err != nil{
		return mapErrors.MapErrorsUser(create_err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("user created"))
	return nil
}