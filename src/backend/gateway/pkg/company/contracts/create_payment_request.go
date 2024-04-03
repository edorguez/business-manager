package contracts

import "github.com/go-playground/validator/v10"

type CreatePaymentRequest struct {
	CompanyId            int64   `json:"companyId" validate:"required"`
	Name                 string  `json:"name" validate:"required,max=50"`
	Bank                 *string `json:"bank" validate:"omitempty,max=50"`
	AccountNumber        *string `json:"accountNumber" validate:"omitempty,max=20"`
	AccountType          *string `json:"accountType" validate:"omitempty,max=20"`
	IdentificationNumber *string `json:"identificationNumber" validate:"omitempty,max=20"`
	IdentificationType   *string `json:"identificationType" validate:"omitempty,max=1"`
	Phone                *string `json:"phone" validate:"omitempty,max=11"`
	Email                *string `json:"email" validate:"omitempty,max=100"`
	PaymentTypeId        int64   `json:"paymentTypeId" validate:"required"`
}

func (c *CreatePaymentRequest) Validate() error {
	validate := validator.New()

	return validate.Struct(c)
}
