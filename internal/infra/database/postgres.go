package database

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadDbConfigs() string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Erro ao carregar .env")
	}

	cred := os.Getenv("DATABASE_URL")

	return cred
}

func Connect() *gorm.DB {
	dsn := LoadDbConfigs()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil{
		log.Fatalf("ocorreu um erro ao se conectar com o banco de dados")
	}

	return db
}