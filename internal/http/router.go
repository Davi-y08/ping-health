package http

import (
	"net/http"
	u "ping-health/internal/application/user"
	m "ping-health/internal/application/monitor"
	h "ping-health/internal/http/handlers"
	"ping-health/internal/http/middlewares"
	r "ping-health/internal/repository"

	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) http.Handler{
	mux := http.NewServeMux()

	userRepo 	:= r.NewUserRepository(db)
	userService := u.NewUserService(userRepo)
	userHandler := h.NewUserHandler(userService)

	monitorRepo := r.NewMonitorRepository(db)
	monitorService := m.NewMonitorService(monitorRepo)
	monitorHandler := h.NewMonitorHandler(monitorService)

	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	mux.HandleFunc("POST /users", middlewares.ErrorsMiddleware(userHandler.CreateUserHandler))
	mux.HandleFunc("POST /users/login", middlewares.ErrorsMiddleware(userHandler.LoginHandler))

	mux.HandleFunc("POST /monitor", middlewares.ErrorsMiddleware(middlewares.JWTAuthMiddleware(
		monitorHandler.CreateMonitorHandler,
	)))
	
	return mux
}