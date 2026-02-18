package middlewares

import (
	"context"
	"net/http"
	"os"
	"ping-health/internal/httpx"

	"github.com/golang-jwt/jwt/v5"
)

var jwt_key []byte

func LoadJwtConfig() {
	jwt_key_aux := os.Getenv("JWT_SECRET_KEY")

	if jwt_key_aux == ""{
		panic("jwt_key não configurada")
	}

	jwt_key = []byte(jwt_key_aux)
}

func JWTAuthMiddleware(next AppHandler) AppHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		cookie, err := r.Cookie("access_token")

		if err != nil{
			return &httpx.AppError{
				Status:  http.StatusUnauthorized,
				Message: "missing authentication token",
				Err:     err,
			}
		}

		claims := &jwt.RegisteredClaims{}

		token, err := jwt.ParseWithClaims(
			cookie.Value,
			claims,
			func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok{
					return nil, jwt.ErrSignatureInvalid
				}

				return jwt_key, nil
			},
			jwt.WithIssuer("api"),
		)

		if err != nil || !token.Valid{
			return &httpx.AppError{
				Status:  http.StatusUnauthorized,
				Message: "invalid or expired token",
				Err:     err,
			}
		}

		ctx := context.WithValue(r.Context(), "user_id", claims.Subject)
		r = r.WithContext(ctx)

		return next(w, r)
	}
}