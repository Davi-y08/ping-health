package http

import (
	"net/http"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) http.Handler{
	mux := http.NewServeMux()

	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	return mux
}