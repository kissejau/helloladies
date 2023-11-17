package config

import (
	"fmt"
	"helloladies/apps/backend/internal/providers/postgres"
	"helloladies/apps/backend/internal/server"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

const (
	ErrFieldNotFound = "field not found in environment variables"
)

type Config struct {
	Server   server.Config
	Postgres postgres.Config
}

func New() (Config, error) {
	var (
		config         Config
		serverConfig   server.Config
		postgresConfig postgres.Config
	)

	wd, err := os.Getwd()
	if err != nil {
		return Config{}, fmt.Errorf("os.Getwd: %w", err)
	}

	envpath := filepath.Join(wd, ".env")

	err = godotenv.Load(envpath)
	if err != nil {
		return Config{}, fmt.Errorf("godotenv.Load: %w", err)
	}

	serverPort := os.Getenv("SERVER_PORT")
	if err := Validate(serverPort, "SERVER_PORT"); err != nil {
		return Config{}, err
	}

	serverConfig = server.Config{
		Port: serverPort,
	}

	postgresHost := os.Getenv("POSTGRES_HOST")
	if err := Validate(postgresHost, "POSTGRES_HOST"); err != nil {
		return Config{}, err
	}

	postgresPort := os.Getenv("POSTGRES_PORT")
	if err := Validate(postgresPort, "POSTGRES_PORT"); err != nil {
		return Config{}, err
	}

	postgresUser := os.Getenv("POSTGRES_USER")
	if err := Validate(postgresUser, "POSTGRES_USER"); err != nil {
		return Config{}, err
	}

	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	if err := Validate(postgresPassword, "POSTGRES_PASSWORD"); err != nil {
		return Config{}, err
	}

	postgresDatabase := os.Getenv("POSTGRES_DATABASE")
	if err := Validate(postgresDatabase, "POSTGRES_DATABASE"); err != nil {
		return Config{}, err
	}

	postgresConfig = postgres.Config{
		Host:     postgresHost,
		Port:     postgresPort,
		User:     postgresUser,
		Password: postgresPassword,
		Database: postgresDatabase,
	}

	config = Config{
		Server:   serverConfig,
		Postgres: postgresConfig,
	}

	return config, nil
}

func Validate(val string, field string) error {
	if len(val) == 0 {
		return fmt.Errorf("%s %s", field, ErrFieldNotFound)
	}
	return nil
}
