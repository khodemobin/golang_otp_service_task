package dto

import (
	"github.com/invopop/validation"
	"github.com/invopop/validation/is"
)

type (
	OTPRequest struct {
		Phone string `json:"phone"`
	}
)

func (req OTPRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Phone, validation.Required, is.Digit, validation.Length(0, 11)),
	)
}
