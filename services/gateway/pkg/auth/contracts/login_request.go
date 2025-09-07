package contracts

import "github.com/go-playground/validator/v10"

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (l *LoginRequest) Validate() error {
	validate := validator.New()

	return validate.Struct(l)
}
