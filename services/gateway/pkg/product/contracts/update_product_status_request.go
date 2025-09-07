package contracts

import "github.com/go-playground/validator/v10"

type UpdateProductStatusRequest struct {
	ProductStatus *uint32 `json:"productStatus" validate:"required"`
}

func (p *UpdateProductStatusRequest) Validate() error {
	validate := validator.New()

	return validate.Struct(p)
}
