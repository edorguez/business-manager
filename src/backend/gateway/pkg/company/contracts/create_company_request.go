package contracts

import "github.com/go-playground/validator/v10"

type CreateCompanyRequest struct {
	Name     string  `json:"name" validate:"required,max=50"`
	ImageUrl *string `json:"imageUrl" validate:"omitempty,required"`
}

func (c *CreateCompanyRequest) Validate() error {
	validate := validator.New()

	return validate.Struct(c)
}
