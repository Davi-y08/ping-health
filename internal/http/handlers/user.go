package handlers

import (
	"net/http"
	service "ping-health/internal/application/user"
)

type UserHandler struct{
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service }
}

func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("user created"))
}