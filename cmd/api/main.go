package main

import (
	"log"
	"net/http"
	httpServer "ping-health/internal/http"
)

func main() {
	router := httpServer.SetupRouter()

	log.Println("servidor rodando na porta :8080")

	err := http.ListenAndServe(":8080", router)

	if err != nil{
		log.Fatal("erro ao inicializar servidor: ", err.Error())
	}
}