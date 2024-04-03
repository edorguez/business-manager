package contracts

import "github.com/go-playground/validator/v10"

type UpdateCustomerRequest struct {
	FirstName            string  `json:"firstName" validate:"required,max=20"`
	LastName             *string `json:"lastName" validate:"omitempty,max=20"`
	Email                *string `json:"email" validate:"omitempty,email,max=100"`
	Phone                *string `json:"phone" validate:"omitempty,max=11"`
	IdentificationNumber string  `json:"identificationNumber" validate:"required,max=20"`
	IdentificationType   string  `json:"identificationType" validate:"required,max=1"`
}

func (c *UpdateCustomerRequest) Validate() error {
	validate := validator.New()

	return validate.Struct(c)
}
