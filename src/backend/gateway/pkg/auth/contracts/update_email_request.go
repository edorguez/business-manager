package contracts

import "github.com/go-playground/validator/v10"

type UpdateEmailRequest struct {
	Email string `json:"email" validate:"required,max=100"`
}

func (c *UpdateEmailRequest) Validate() error {
	validate := validator.New()

	return validate.Struct(c)
}
