package service

import (
	"github.com/gin-gonic/gin"

	"github.com/ch3nnn/webstack-go/internal/dal/repository"
	"github.com/ch3nnn/webstack-go/pkg/jwt"
	"github.com/ch3nnn/webstack-go/pkg/log"
)

type Service struct {
	Logger     *log.Logger
	Jwt        *jwt.JWT
	SvcCtx     *gin.Engine
	Repository *repository.Repository
}

func NewService(
	svcCtx *gin.Engine,
	logger *log.Logger,
	jwt *jwt.JWT,
	repository *repository.Repository,
) *Service {
	return &Service{
		Logger:     logger,
		Jwt:        jwt,
		SvcCtx:     svcCtx,
		Repository: repository,
	}
}
