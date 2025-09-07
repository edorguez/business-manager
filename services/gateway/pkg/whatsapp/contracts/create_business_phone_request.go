package contracts

import "github.com/go-playground/validator"

type CreateBusinessPhoneRequest struct {
	CompanyId int64  `json:"companyId" validate:"required"`
	Phone     string `json:"phone" validate:"required,max=11"`
}

func (c *CreateBusinessPhoneRequest) Validate() error {
	validate := validator.New()

	return validate.Struct(c)
}
