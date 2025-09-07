package contracts

import "github.com/go-playground/validator/v10"

type UpdateBusinessPhoneRequest struct {
	CompanyId int64  `json:"companyId" validate:"required"`
	Phone     string `json:"phone" validate:"omitempty,max=11"`
}

func (c *UpdateBusinessPhoneRequest) Validate() error {
	validate := validator.New()

	return validate.Struct(c)
}
