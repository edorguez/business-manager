package contracts

import "github.com/go-playground/validator/v10"

type UpdateProductRequest struct {
	Name        string   `json:"name" validate:"required,max=50"`
	Description *string  `json:"description" validate:"omitempty,max=50"`
	Sku         *string  `json:"sku" validate:"omitempty,max=12"`
	Quantity    uint64   `json:"quantity" validate:"required"`
	Price       uint64   `json:"price" validate:"required"`
	Images      []string `json:"images"`
}

func (p *UpdateProductRequest) Validate() error {
	validate := validator.New()

	return validate.Struct(p)
}
