package handlers

import (
	"Sunat/internal/services"
	"go.uber.org/zap"
)

func NewHandler(service *services.Services, logger *zap.Logger) *Handler {
	return &Handler{Service: service, Logger: logger}
}
