package app

import (
	"log"
	"users-services/internal/controller/http"
	"users-services/internal/repo"
	"users-services/pkg/postgres"
)

func Run() {
	log.Println("Users-services start work...")

	// Repository
	pg := postgres.NewPostgres()
	dbRepo := repo.New(pg)
	dbRepo.MigrateModel()

	// Controller
	c := http.New(dbRepo)
	c.NewRoute()
	err := c.Route.Run(":8080")
	if err != nil {
		log.Fatal("Error, gin run: ", err)
	}
}
