package repository

import "github.com/jmoiron/sqlx"

type AuthRepositoryImpl struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) AuthRepositoryImpl {
	return AuthRepositoryImpl{
		db: db,
	}
}
