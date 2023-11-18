package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func New(cfg Config) (*sql.DB, error) {
	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Database,
		cfg.Password)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("db.Ping: %w", err)
	}

	return db, nil
}
