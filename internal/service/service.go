package service

import (
	"github.com/ch3nnn/webstack-go/internal/dal/repository"
	"github.com/ch3nnn/webstack-go/pkg/jwt"
	"github.com/ch3nnn/webstack-go/pkg/log"
)

type Service struct {
	Logger     *log.Logger
	Jwt        *jwt.JWT
	Repository *repository.Repository
}

func NewService(
	logger *log.Logger,
	jwt *jwt.JWT,
	repository *repository.Repository,
) *Service {
	return &Service{
		Logger:     logger,
		Jwt:        jwt,
		Repository: repository,
	}
}
