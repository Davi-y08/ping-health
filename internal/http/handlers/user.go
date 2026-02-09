package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	appUser "ping-health/internal/application/user"
	service "ping-health/internal/application/user"
	mapErrors "ping-health/internal/http/http_errors"
	"ping-health/internal/httpx"
	security "ping-health/internal/infra/security"
	"github.com/google/uuid"
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

func (h *UserHandler) LoginHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		return httpx.MethodNotAllowed(errors.New("method not allowed"))
	}

	var dto appUser.LoginDto

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil{
		return httpx.BadRequest(errors.New("corpo inválido"))
	}

	user, err := h.service.Login(r.Context(), dto)

	if err != nil{
		return mapErrors.MapErrorsUser(err)
	}

	token, err := security.GenerateTokenJWT(user.ID)

	csfr := uuid.NewString()

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    token,
		Path:     "/",
		MaxAge:	  (24 * 1) / 2,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	})

	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(map[string]string{
		"message": "login realizado com sucesso!",
		"csfr": csfr,
	})
}