package contracts

import "github.com/go-playground/validator/v10"

type CreateOrderProductRequest struct {
	ProductId string `json:"productId" validate:"required"`
	Quantity  uint64 `json:"quantity" validate:"required"`
	Price     uint64 `json:"price" validate:"required"`
}

func (p *CreateOrderProductRequest) Validate() error {
	validate := validator.New()

	return validate.Struct(p)
}
