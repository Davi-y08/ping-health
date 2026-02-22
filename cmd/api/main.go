package main

import (
	"context"
	"log"
	"net/http"
	"ping-health/internal/application/monitor"
	httpServer "ping-health/internal/http"
	db "ping-health/internal/infra/database"
	repo "ping-health/internal/repository"
)

func main() {
	ctx := context.Background()
	database := db.Connect()
	router := httpServer.SetupRouter(database)
	db.RunMigrations(database)

	monitorRepo := repo.NewMonitorRepository(database)
	service := monitor.NewMonitorService(monitorRepo)

	if err := service.StartAllMonitors(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("servidor rodando na porta :8080")

	err := http.ListenAndServe(":8080", router)

	if err != nil{
		log.Fatal("erro ao inicializar servidor: ", err.Error())
	}
}