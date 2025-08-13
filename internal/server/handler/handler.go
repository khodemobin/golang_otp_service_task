package handler

import (
	"github.com/khodemobin/golang_otp_service_task/internal/service"
	"github.com/khodemobin/golang_otp_service_task/pkg/logger"
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
