package contracts

import "github.com/go-playground/validator/v10"

type UpdatePasswordRequest struct {
	Password string `json:"password" validate:"required,max=20"`
}

func (c *UpdatePasswordRequest) Validate() error {
	validate := validator.New()

	return validate.Struct(c)
}
