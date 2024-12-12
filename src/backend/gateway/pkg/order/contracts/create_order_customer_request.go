package contracts

import "github.com/go-playground/validator/v10"

type CreateOrderCustomerRequest struct {
	FirstName            string  `json:"firstName" validate:"required,max=20"`
	LastName             *string `json:"lastName" validate:"omitempty,max=20"`
	Phone                *string `json:"phone" validate:"omitempty,max=11"`
	IdentificationNumber string  `json:"identificationNumber" validate:"required,max=20"`
	IdentificationType   string  `json:"identificationType" validate:"required,max=1"`
}

func (c *CreateOrderCustomerRequest) Validate() error {
	validate := validator.New()

	return validate.Struct(c)
}
