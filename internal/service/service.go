package service

import (
	"github.com/khodemobin/golang_boilerplate/internal/config"
	"github.com/khodemobin/golang_boilerplate/internal/server/dto"
	"github.com/khodemobin/golang_boilerplate/pkg/cache"
	"github.com/khodemobin/golang_boilerplate/pkg/logger"
	"gorm.io/gorm"
)

type IOTPService interface {
	Send(*dto.OTPRequest) error
	Verify(*dto.OTPVerifyRequest) (*dto.OTPVerifyResponse, error)
}

type IJWTService interface {
	GenerateJWT(userID uint, phone string) (string, error)
	ValidateJWT(tokenString string) (*JWTClaims, error)
}

type IAuthService interface {
	Login(phone string) (string, *dto.User, error)
}

type Service struct {
	OTPService  IOTPService
	JWTService  IJWTService
	AuthService IAuthService
}

func NewService(cfg *config.Config, db *gorm.DB, logger logger.Logger, cache cache.Cache) *Service {
	jwtService := newJWTService(cfg)
	authService := newAuthService(db, jwtService)
	otpService := newOTPService(logger, cache, authService)

	return &Service{
		OTPService:  otpService,
		JWTService:  jwtService,
		AuthService: authService,
	}
}
