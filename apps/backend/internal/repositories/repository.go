package repositories

import (
	"database/sql"

	"github.com/sirupsen/logrus"
)

type Repositories struct {
}

func NewRepositories(db *sql.DB, log *logrus.Logger) *Repositories {
	return &Repositories{}
}
