package main

import (
	"log"
	"net/http"
	httpServer "ping-health/internal/http"
	db "ping-health/internal/infra/database"
)

func main() {
	database := db.Connect()
	router := httpServer.SetupRouter(database)
	db.RunMigrations(database)
	log.Println("servidor rodando na porta :8080")

	err := http.ListenAndServe(":8080", router)

	if err != nil{
		log.Fatal("erro ao inicializar servidor: ", err.Error())
	}

}