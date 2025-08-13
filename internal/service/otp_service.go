package service

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/khodemobin/golang_boilerplate/internal/server/dto"
	"github.com/khodemobin/golang_boilerplate/pkg/apperror"
	"github.com/khodemobin/golang_boilerplate/pkg/cache"
	"github.com/khodemobin/golang_boilerplate/pkg/logger"
)

type IOTPServiceImpl struct {
	logger      logger.Logger
	cache       cache.Cache
	authService IAuthService
}

type OTPData struct {
	OTP       string    `json:"otp"`
	Phone     string    `json:"phone"`
	ExpiresAt time.Time `json:"expires_at"`
	Attempts  int       `json:"attempts"`
}

func newOTPService(logger logger.Logger, cache cache.Cache, authService IAuthService) IOTPService {
	return &IOTPServiceImpl{
		logger:      logger,
		cache:       cache,
		authService: authService,
	}
}

func (s *IOTPServiceImpl) Send(req *dto.OTPRequest) error {
	otp, err := s.generateOTP()
	if err != nil {
		return err
	}

	otpData := OTPData{
		OTP:       otp,
		Phone:     req.Phone,
		ExpiresAt: time.Now().Add(2 * time.Minute),
		Attempts:  0,
	}

	if err := s.cache.Set(s.otpCacheKey(req.Phone), otpData, 2*time.Minute); err != nil {
		return err
	}

	s.logger.Info("Phone Number : " + req.Phone)
	s.logger.Info("OPT Code : " + otp)
	s.logger.Info("Expired : " + otpData.ExpiresAt.Format("15:04:05"))
	s.logger.Info("================================")

	return nil
}

func (s *IOTPServiceImpl) Verify(req *dto.OTPVerifyRequest) (*dto.OTPVerifyResponse, error) {
	cacheKey := s.otpCacheKey(req.Phone)

	var otpData OTPData
	if err := s.cache.Get(cacheKey, &otpData); err != nil {
		return nil, apperror.BadRequest(fmt.Errorf("otp not found"))
	}

	if otpData.Attempts >= 3 {
		if err := s.cache.Delete(cacheKey); err != nil {
			return nil, err
		}
		return nil, apperror.BadRequest(fmt.Errorf("too many attempts"))
	}

	otpData.Attempts++
	if err := s.cache.Set(cacheKey, otpData, 2*time.Minute); err != nil {
		return nil, err
	}

	if otpData.OTP != req.OTP {
		return nil, apperror.BadRequest(fmt.Errorf("otp not found"))
	}

	if err := s.cache.Delete(cacheKey); err != nil {
		return nil, err
	}

	token, user, err := s.authService.Login(req.Phone)
	if err != nil {
		return nil, err
	}

	return &dto.OTPVerifyResponse{
		Token: token,
		User:  user,
	}, nil
}

func (s *IOTPServiceImpl) generateOTP() (string, error) {
	maxVal := big.NewInt(999999)
	minVal := big.NewInt(100000)

	randomNum, err := rand.Int(rand.Reader, maxVal.Sub(maxVal, minVal).Add(maxVal, big.NewInt(1)))
	if err != nil {
		return "", err
	}

	otp := randomNum.Add(randomNum, minVal)
	return otp.String(), nil
}

func (s *IOTPServiceImpl) otpCacheKey(phone string) string {
	return fmt.Sprintf("otp:%s", phone)
}
