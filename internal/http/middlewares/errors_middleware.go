package middlewares

import (
	"encoding/json"
	"errors"
	"net/http"
	"ping-health/internal/httpx"
)

type AppHandler func(w http.ResponseWriter, r *http.Request) error

func ErrorsMiddleware(next AppHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := next(w, r)
		if err == nil{
			return
		}

		var appErr *httpx.AppError
		if errors.As(err, &appErr) {
			w.WriteHeader(appErr.Status)
			_ = json.NewEncoder(w).Encode(map[string]string{
				"error": appErr.Message,
			})

			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"error": "internal error",
		})
	}
}