package contracts

import "github.com/go-playground/validator/v10"

type CreateOrderProductRequest struct {
}

func (p *CreateOrderProductRequest) Validate() error {
	validate := validator.New()

	return validate.Struct(p)
}
