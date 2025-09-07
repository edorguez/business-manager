package contracts

import "github.com/go-playground/validator/v10"

type CreateUserRequest struct {
	CompanyId int64  `json:"companyId" validate:"required"`
	RoleId    int64  `json:"roleId" validate:"required"`
	Email     string `json:"email" validate:"required,max=100"`
	Password  string `json:"password" validate:"required,max=20"`
}

func (c *CreateUserRequest) Validate() error {
	validate := validator.New()

	return validate.Struct(c)
}
