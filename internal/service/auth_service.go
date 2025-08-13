package service

import (
	"errors"

	"github.com/khodemobin/golang_boilerplate/internal/model"
	"github.com/khodemobin/golang_boilerplate/internal/server/dto"
	"gorm.io/gorm"
)

type IAuthServiceImpl struct {
	db         *gorm.DB
	jwtService IJWTService
}

func newAuthService(db *gorm.DB, jwtService IJWTService) IAuthService {
	return &IAuthServiceImpl{
		db:         db,
		jwtService: jwtService,
	}

}

func (s *IAuthServiceImpl) Login(phone string) (string, *dto.User, error) {
	var user model.User
	err := s.db.Where("phone = ?", phone).First(&user).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		user.Phone = phone
		if err := s.db.Create(&user).Error; err != nil {
			return "", nil, err
		}
	}

	token, err := s.jwtService.GenerateJWT(user.ID, user.Phone)
	if err != nil {
		return "", nil, err
	}

	return token, &dto.User{
		ID:    user.ID,
		Phone: user.Phone,
	}, nil
}
