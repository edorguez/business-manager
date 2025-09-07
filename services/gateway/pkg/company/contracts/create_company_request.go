package contracts

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type CreateCompanyRequest struct {
	Name            string    `json:"name" validate:"required,max=50"`
	LastPaymentDate time.Time `json:"lastPaymentDate" validate:"required"`
}

func (c *CreateCompanyRequest) Validate() error {
	validate := validator.New()

	return validate.Struct(c)
}
