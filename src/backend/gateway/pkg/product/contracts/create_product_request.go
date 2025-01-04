package contracts

import "github.com/go-playground/validator/v10"

type CreateProductRequest struct {
	CompanyId     int64   `json:"companyId" validate:"required"`
	Name          string  `json:"name" validate:"required,max=50"`
	Description   *string `json:"description" validate:"omitempty,max=50"`
	Sku           *string `json:"sku" validate:"omitempty,max=12"`
	Quantity      uint64  `json:"quantity" validate:"required"`
	Price         uint64  `json:"price" validate:"required"`
	ProductStatus uint32  `json:"productStatus" validate:"required"`
}

func (p *CreateProductRequest) Validate() error {
	validate := validator.New()

	return validate.Struct(p)
}
