package handler

import (
	"github.com/khodemobin/golang_boilerplate/internal/service"
	"github.com/khodemobin/golang_boilerplate/pkg/logger"
)

type Handler struct {
	AuthHandler *AuthHandler
	UserHandler *UserHandler
}

func NewHandler(logger logger.Logger, service *service.Service) *Handler {
	return &Handler{
		AuthHandler: newAuthHandler(logger, service.OTPService),
		UserHandler: newUserHandler(logger, service.UserService),
	}
}
