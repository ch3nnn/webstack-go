package handler

import (
	"github.com/ch3nnn/webstack-go/pkg/log"
)

type Handler struct {
	Logger *log.Logger
}

func NewHandler(logger *log.Logger) *Handler {
	return &Handler{
		Logger: logger,
	}
}
