package contracts

import "github.com/go-playground/validator/v10"

type CreateCustomerRequest struct {
	CompanyId            int64   `json:"companyId" validate:"required"`
	FirstName            string  `json:"firstName" validate:"required,max=20"`
	LastName             *string `json:"lastName" validate:"omitempty,max=20"`
	Email                *string `json:"email" validate:"omitempty,max=100"`
	Phone                *string `json:"phone" validate:"omitempty,max=11"`
	IdentificationNumber string  `json:"identificationNumber" validate:"required,max=9"`
	IdentificationType   string  `json:"identificationType" validate:"required,max=1"`
}

func (c *CreateCustomerRequest) Validate() error {
	validate := validator.New()

	return validate.Struct(c)
}
