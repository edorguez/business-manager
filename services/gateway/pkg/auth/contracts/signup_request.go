package contracts

import "github.com/go-playground/validator/v10"

type SignUpRequest struct {
	Company SignUpCompany `json:"company" validate:"required"`
	User    SignUpUser    `json:"user" validate:"required"`
}

type SignUpCompany struct {
	Name          string `json:"name" validate:"required"`
	NameFormatUrl string `json:"nameFormatUrl" validate:"required"`
	Phone         string `json:"phone" validate:"required"`
}

type SignUpUser struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (s *SignUpRequest) Validate() error {
	validate := validator.New()

	return validate.Struct(s)
}
