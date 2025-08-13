package dto

import "github.com/invopop/validation"

type (
	OTPRequest struct {
		Phone string `json:"phone"`
	}
)

func (req OTPRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Phone, validation.Required, validation.Length(0, 11)),
	)
}
