package security

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

var jwt_key []byte

func LoadJWTConfig() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Erro ao carregar .env")
	}

	byte_array := os.Getenv("JWT_KEY")

	jwt_key = []byte(byte_array)
}

func GenerateTokenJWT(user_id uuid.UUID) (string, error){
	claims := jwt.RegisteredClaims{
		Subject: user_id.String(),
		Issuer: "api",
		IssuedAt: jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token_string, err := token.SignedString(jwt_key)

	if err != nil {
		return "", errors.New("erro ao gerar token jwt")
	}

	return token_string, nil
}