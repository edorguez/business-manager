package contracts

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
)

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

	if err := validate.Struct(s); err != nil {
		return err
	}

	reserved := os.Getenv("RESERVED_SUBDOMAINS")
	if reserved == "" {
		reserved = "www,tourism-agency,gps-medical-assistance,business-manager"
	}

	for _, item := range strings.Split(reserved, ",") {
		if strings.EqualFold(strings.TrimSpace(item), s.Company.NameFormatUrl) {
			return fmt.Errorf("subdomain %q is reserved", s.Company.NameFormatUrl)
		}
	}

	return nil
}
