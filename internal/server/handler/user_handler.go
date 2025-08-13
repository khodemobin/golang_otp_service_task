package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khodemobin/golang_boilerplate/internal/server/dto"
	"github.com/khodemobin/golang_boilerplate/internal/service"
	"github.com/khodemobin/golang_boilerplate/pkg/apperror"
	"github.com/khodemobin/golang_boilerplate/pkg/logger"
	"github.com/khodemobin/golang_boilerplate/pkg/response"
)

type UserHandler struct {
	logger      logger.Logger
	userService service.IUserService
}

func newUserHandler(logger logger.Logger, userService service.IUserService) *UserHandler {
	return &UserHandler{
		logger:      logger,
		userService: userService,
	}
}

// Index list of users
//
//	@Summary      get list of users
//	@Accept       json
//	@Produce      json
//
// @param Authorization header string true "Authorization"
// @Param	page	query	int	false	"page"
// @Param	count	query	int	false	"count"
// @Param	phone	query	string	false	"phone"
//
// @Success 200 {object} dto.UserListResponse
// @Router       /users [get]
func (h *UserHandler) Index(c *fiber.Ctx) error {
	req := new(dto.UserListRequest)
	if err := c.QueryParser(req); err != nil {
		return response.ErrorBuilder().FromError(apperror.BadRequest(err)).Send(c)
	}

	if err := req.Validate(); err != nil {
		return response.ErrorBuilder().FromError(apperror.Validation(err)).Send(c)
	}

	result, err := h.userService.Index(req)
	if err != nil {
		h.logger.Error(err)
		return response.ErrorBuilder().FromError(err).Send(c)
	}

	return response.SuccessBuilder().WithData(result.Users).Send(c)
}

// Get user
//
//	@Summary      get user by id
//	@Accept       json
//	@Produce      json
//
// @param Authorization header string true "Authorization"
// @Param        id   path  int  true  "User ID"
//
// @Success 200 {object} dto.UserGetResponse
// @Router       /users/{id} [get]
func (h *UserHandler) Get(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return response.ErrorBuilder().FromError(err).Send(c)
	}

	result, err := h.userService.Find(uint(id))
	if err != nil {
		return response.ErrorBuilder().FromError(err).Send(c)
	}

	return response.SuccessBuilder().WithData(result).Send(c)
}
