package http

import (
	"net/http"
	"gorm.io/gorm"
	r "ping-health/internal/repository"
	u "ping-health/internal/application/user"
	h "ping-health/internal/http/handlers"
)

func SetupRouter(db *gorm.DB) http.Handler{
	mux := http.NewServeMux()

	userRepo 	:= r.NewUserRepository(db)
	userService := u.NewUserService(userRepo)
	userHandler := h.NewUserHandler(userService)

	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	mux.HandleFunc("POST /users", userHandler.CreateUserHandler)

	return mux
}