package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khodemobin/golang_boilerplate/internal/server/dto"
	"github.com/khodemobin/golang_boilerplate/internal/service"
	"github.com/khodemobin/golang_boilerplate/pkg/apperror"
	"github.com/khodemobin/golang_boilerplate/pkg/logger"
	"github.com/khodemobin/golang_boilerplate/pkg/response"
)

type AuthHandler struct {
	logger     logger.Logger
	otpService service.IOTPService
}

func newAuthHandler(logger logger.Logger, otpService service.IOTPService) *AuthHandler {
	return &AuthHandler{
		logger:     logger,
		otpService: otpService,
	}
}

// SendOTP send login/register otp code
//
//	@Summary      Send login/register otp code
//	@Accept       json
//	@Produce      json
//
// @Param request body dto.OTPRequest true "query params"
//
//	@Success      200
//	@Router       /auth/otp/send [post]
func (h *AuthHandler) SendOTP(c *fiber.Ctx) error {
	req := new(dto.OTPRequest)
	if err := c.BodyParser(&req); err != nil {
		return response.ErrorBuilder().FromError(apperror.BadRequest(err)).Send(c)
	}

	if err := req.Validate(); err != nil {
		return response.ErrorBuilder().FromError(apperror.Validation(err)).Send(c)
	}

	if err := h.otpService.Send(req); err != nil {
		h.logger.Error(err)
		return response.ErrorBuilder().FromError(err).Send(c)
	}

	return response.SuccessBuilder().WithData(req.Phone).WithMessage("OTP Send").Send(c)
}

// VerifyOTP verify login/register otp code
//
//	@Summary      Very login/register otp code
//	@Description  If user exists in database return new token and if not create new one.
//	@Accept       json
//	@Produce      json
//
// @Param request body dto.OTPVerifyRequest true "query params"
//
// @Success 200 {object} dto.OTPVerifyResponse
// @Router       /auth/otp/verify [post]
func (h *AuthHandler) VerifyOTP(c *fiber.Ctx) error {
	req := new(dto.OTPVerifyRequest)
	if err := c.BodyParser(&req); err != nil {
		return response.ErrorBuilder().FromError(apperror.BadRequest(err)).Send(c)
	}

	if err := req.Validate(); err != nil {
		return response.ErrorBuilder().FromError(apperror.Validation(err)).Send(c)
	}

	result, err := h.otpService.Verify(req)
	if err != nil {
		h.logger.Error(err)
		return response.ErrorBuilder().FromError(err).Send(c)
	}

	return response.SuccessBuilder().WithData(result).Send(c)
}
