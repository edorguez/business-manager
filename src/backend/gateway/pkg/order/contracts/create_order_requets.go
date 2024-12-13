package contracts

import "github.com/go-playground/validator/v10"

type CreateOrderRequest struct {
	CompanyId int64                       `json:"companyId" validate:"required"`
	Customer  CreateOrderCustomerRequest  `json:"customer" validate:"required"`
	Products  []CreateOrderProductRequest `json:"products" validate:"required"`
}

func (p *CreateOrderRequest) Validate() error {
	validate := validator.New()

	return validate.Struct(p)
}
