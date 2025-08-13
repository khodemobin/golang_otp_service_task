package service

import (
	"errors"

	"github.com/khodemobin/golang_otp_service_task/internal/model"
	"github.com/khodemobin/golang_otp_service_task/internal/server/dto"
	"github.com/khodemobin/golang_otp_service_task/pkg/apperror"
	"github.com/khodemobin/golang_otp_service_task/pkg/pgsql/scope"
	"gorm.io/gorm"
)

type IUserServiceImpl struct {
	db *gorm.DB
}

func newUserService(db *gorm.DB) IUserService {
	return &IUserServiceImpl{
		db: db,
	}
}

func (s IUserServiceImpl) Index(req *dto.UserListRequest) (*dto.UserListResponse, error) {
	var users []model.User

	query := s.db.Scopes(scope.Paginate(&req.PaginationRequest))
	if req.Phone != nil {
		query.Where("phone LIKE ?", "%"+*req.Phone+"%")
	}

	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}

	mappedUser := make([]dto.User, len(users))
	for i, user := range users {
		mappedUser[i] = dto.User{
			ID:    user.ID,
			Phone: user.Phone,
		}
	}

	return &dto.UserListResponse{
		Users: mappedUser,
	}, nil
}

func (s IUserServiceImpl) Find(id uint) (*dto.UserGetResponse, error) {
	var user model.User
	err := s.db.Where("id = ?", id).First(&user).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, apperror.NotFound(errors.New("user not found"))
	}

	return &dto.UserGetResponse{
		User: &dto.User{
			ID:    user.ID,
			Phone: user.Phone,
		},
	}, nil
}
