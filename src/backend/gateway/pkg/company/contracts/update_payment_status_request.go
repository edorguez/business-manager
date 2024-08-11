package contracts

import "github.com/go-playground/validator/v10"

type UpdatePaymentStatusRequest struct {
	Status bool `json:"status" validate:"required"`
}

func (c *UpdatePaymentStatusRequest) Validate() error {
	validate := validator.New()

	return validate.Struct(c)
}
