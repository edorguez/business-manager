package db

import (
	"testing"

	"github.com/EdoRguez/business-manager/client-svc/pkg/util"
)

func createRandomClient(t *testing.T) ClientClient {
	arg := CreateClientParams{
		CompanyID:            0,
		FirstName:            util.RandomName(),
		LastName:             util.NewSqlNullString(util.RandomName()),
		Email:                util.NewSqlNullString(util.RandomEmail()),
		Phone:                util.NewSqlNullString(util.RandomPhoneNumber()),
		IdentificationNumber: util.RandomIndentificationNumber(),
		IdentificationType:   "",
	}

}
