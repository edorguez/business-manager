package contracts

import "github.com/go-playground/validator/v10"

type UpdateCompanyRequest struct {
	Name          string `json:"name" validate:"required,max=50"`
	NameFormatUrl string `json:"nameFormatUrl" validate:"required"`
}

func (c *UpdateCompanyRequest) Validate() error {
	validate := validator.New()

	return validate.Struct(c)
}
