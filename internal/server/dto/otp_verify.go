package dto

import "github.com/invopop/validation"

type (
	OTPVerifyRequest struct {
		Phone string `json:"phone" validate:"required"`
		OTP   string `json:"otp" validate:"required,len=6"`
	}

	OTPVerifyResponse struct {
		Token string `json:"token"`
		User  *User  `json:"user"`
	}
)

func (req OTPVerifyRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Phone, validation.Required, validation.Length(0, 11)),
		validation.Field(&req.OTP, validation.Required, validation.Length(0, 6)),
	)
}
