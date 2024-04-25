package contracts

import "github.com/go-playground/validator/v10"

type UpdateProductRequest struct {
	Name          string   `json:"name" validate:"required,max=50"`
	Description   *string  `json:"description" validate:"omitempty,max=50"`
	Sku           *string  `json:"sku" validate:"omitempty,max=12"`
	Price         uint64   `json:"price" validate:"required"`
	Images        []string `json:"images" validate:"required,max=5,dive"`
	ProductStatus uint32   `json:"productStatus" validate:"required"`
}

func (p *UpdateProductRequest) Validate() error {
	validate := validator.New()

	return validate.Struct(p)
}
