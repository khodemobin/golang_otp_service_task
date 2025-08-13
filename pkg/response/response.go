package response

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/khodemobin/golang_boilerplate/pkg/apperror"
)

type FailedResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type FailedResponseBuilder struct {
	response FailedResponse
}

func ErrorBuilder() *FailedResponseBuilder {
	return &FailedResponseBuilder{
		response: FailedResponse{
			Code:    http.StatusInternalServerError,
			Message: "",
		},
	}
}

func (b *FailedResponseBuilder) WithCode(code int) *FailedResponseBuilder {
	b.response.Code = code
	return b
}

func (b *FailedResponseBuilder) WithMessage(message string) *FailedResponseBuilder {
	b.response.Message = message
	return b
}

func (b *FailedResponseBuilder) WithData(data any) *FailedResponseBuilder {
	b.response.Data = data
	return b
}

func (b *FailedResponseBuilder) FromError(err error) *FailedResponseBuilder {
	var appErr *apperror.AppError
	if errors.As(err, &appErr) {
		var ae *apperror.AppError
		errors.As(err, &ae)
		b.response.Code = ae.Code
		b.response.Message = ae.Error()
	} else {
		b.response.Message = err.Error()
	}
	return b
}

func (b *FailedResponseBuilder) Build() FailedResponse {
	return b.response
}

func (b *FailedResponseBuilder) Send(app *fiber.Ctx) error {
	return app.Status(b.response.Code).JSON(b.response)
}

type SuccessResponse struct {
	Data    any    `json:"data"`
	Message string `json:"message"`
	Meta    any    `json:"meta,omitempty"`
}

type SuccessResponseBuilder struct {
	response SuccessResponse
}

func SuccessBuilder() *SuccessResponseBuilder {
	return &SuccessResponseBuilder{
		response: SuccessResponse{
			Data:    nil,
			Message: "",
			Meta:    nil,
		},
	}
}

func (b *SuccessResponseBuilder) WithData(data any) *SuccessResponseBuilder {
	b.response.Data = data
	return b
}

func (b *SuccessResponseBuilder) WithMessage(message string) *SuccessResponseBuilder {
	b.response.Message = message
	return b
}

func (b *SuccessResponseBuilder) WithMeta(meta any) *SuccessResponseBuilder {
	b.response.Meta = meta
	return b
}

func (b *SuccessResponseBuilder) Build() SuccessResponse {
	return b.response
}

func (b *SuccessResponseBuilder) Send(app *fiber.Ctx) error {
	return app.JSON(b.response)
}
