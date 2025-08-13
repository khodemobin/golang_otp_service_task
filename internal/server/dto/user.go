package dto

import "github.com/invopop/validation"

type (
	UserListRequest struct {
		PaginationRequest
		Phone *string `query:"phone"`
	}

	UserListResponse struct {
		Users []User
	}

	UserGetResponse struct {
		*User
	}
)

func (req UserListRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Page, validation.Min(1), validation.Max(100)),
		validation.Field(&req.Count, validation.Min(1), validation.Max(100)),
		validation.Field(&req.Phone, validation.Length(0, 11)),
	)
}
