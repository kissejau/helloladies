package main

import (
	"helloladies/apps/backend/internal/config"
	"helloladies/apps/backend/internal/handlers"
	"helloladies/apps/backend/internal/providers/postgres"
	"helloladies/apps/backend/internal/repositories"
	"helloladies/apps/backend/internal/server"
	"helloladies/apps/backend/internal/services"

	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()

	cfg, err := config.New()
	if err != nil {
		log.Fatalf("config.New: %s", err.Error())
	}
	log.Println("config is loaded")

	db, err := postgres.New(cfg.Postgres)
	if err != nil {
		log.Fatalf("postgres.New: %s", err.Error())
	}
	log.Println("db connection is configured")

	if err := db.Ping(); err != nil {
		log.Fatalf("db.Ping: %s", err.Error())
	}

	repos := repositories.NewRepositories(db, log)

	services := services.New(repos, log)

	handlers := handlers.New(services, log)

	srv := server.New(cfg.Server, log, handlers.InitRoutes())
	if err := srv.Run(); err != nil {
		log.Fatalf("srv.Run: %s", err.Error())
	}
}
