package main

import (
	"helloladies/internal/config"
	"helloladies/internal/handlers"
	postgresProvider "helloladies/internal/providers/postgres"
	"helloladies/internal/repository"
	"helloladies/internal/server"
	"helloladies/internal/service"
	"log"

	_ "helloladies/docs"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

//	@title						helloladies API
//	@version					1.0
//	@description				helloladies API [go + gin]
//
//	@host						localhost:8081
//	@BasePath					/api
//
//	@securityDefinitions.apikey	Token
//	@in							header
//	@name						Authorization
func main() {
	log := logrus.New()
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("config.New: %s", err.Error())
	}
	log.Println("config is loaded")

	db, err := postgresProvider.New(cfg.Postgres)
	if err != nil {
		log.Fatalf("postgres.New: %s", err.Error())
	}
	log.Println("db connection is configured")

	UpMigrations(db)

	if err := db.Ping(); err != nil {
		log.Fatalf("db.Ping: %s", err.Error())
	}

	repos := repository.NewRepositories(db, log)

	services := service.New(repos, log, cfg.JWTConfig)

	handlers := handlers.New(services, cfg.JWTConfig, log)

	srv := server.New(cfg.Server, log, handlers.InitRoutes())
	if err := srv.Run(); err != nil {
		log.Fatalf("srv.Run: %s", err.Error())
	}
}

func UpMigrations(db *sqlx.DB) {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		log.Fatalf("postgres.WithInstance: %s", err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance("file://db/migrations", "postgres", driver)
	if err != nil {
		log.Fatalf("migrate.NewWithDatabaseInstance: %s", err.Error())
	}

	if err := m.Up(); err != nil {
		log.Printf("m.Up: %s", err.Error())
	} else {
		log.Println("Database migration was run successfully")
	}
}
