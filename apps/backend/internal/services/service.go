package services

import (
	"helloladies/apps/backend/internal/repositories"

	"github.com/sirupsen/logrus"
)

type Service struct {
}

func New(repos *repositories.Repositories, log *logrus.Logger) *Service {
	return &Service{}
}
