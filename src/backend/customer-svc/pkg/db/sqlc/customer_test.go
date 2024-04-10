package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/EdoRguez/business-manager/customer-svc/pkg/util/random"
	"github.com/EdoRguez/business-manager/customer-svc/pkg/util/type_converter"
	"github.com/stretchr/testify/require"
)

func createRandomCustomer(t *testing.T) CustomerCustomer {
	randomName := random.Name()
	randomEmail := random.Email()
	randomPhoneNumber := random.PhoneNumber()

	arg := CreateCustomerParams{
		CompanyID:            1,
		FirstName:            random.Name(),
		LastName:             type_converter.NewSqlNullString(&randomName),
		Email:                type_converter.NewSqlNullString(&randomEmail),
		Phone:                type_converter.NewSqlNullString(&randomPhoneNumber),
		IdentificationNumber: random.IndentificationNumber(),
		IdentificationType:   random.IdentificationType(),
	}

	client, err := testQueries.CreateCustomer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, client)

	require.Equal(t, arg.FirstName, client.FirstName)
	require.Equal(t, arg.LastName, client.LastName)
	require.Equal(t, arg.Email, client.Email)
	require.Equal(t, arg.Phone, client.Phone)
	require.Equal(t, arg.IdentificationNumber, client.IdentificationNumber)
	require.Equal(t, arg.IdentificationType, client.IdentificationType)

	require.NotZero(t, client.CompanyID)
	require.NotZero(t, client.CreatedAt)

	return client
}

func TestCreateCustomer(t *testing.T) {
	createRandomCustomer(t)
}

func TestGetCustomer(t *testing.T) {
	client1 := createRandomCustomer(t)
	client2, err := testQueries.GetCustomer(context.Background(), client1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, client2)

	require.Equal(t, client1.ID, client2.ID)
	require.Equal(t, client1.CompanyID, client2.CompanyID)
	require.Equal(t, client1.FirstName, client2.FirstName)
	require.Equal(t, client1.LastName, client2.LastName)
	require.Equal(t, client1.Email, client2.Email)
	require.Equal(t, client1.Phone, client2.Phone)
	require.Equal(t, client1.IdentificationNumber, client2.IdentificationNumber)
	require.Equal(t, client1.IdentificationType, client2.IdentificationType)
	require.WithinDuration(t, client1.CreatedAt, client2.CreatedAt, time.Second)
}

func TestGetCustomers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomCustomer(t)
	}

	arg := GetCustomersParams{
		CompanyID: 0,
		Limit:     5,
		Offset:    5,
	}

	clients, err := testQueries.GetCustomers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, clients, 5)

	for _, client := range clients {
		require.NotEmpty(t, client)
	}
}

func TestUpdateCustomer(t *testing.T) {
	client1 := createRandomCustomer(t)

	arg := UpdateCustomerParams{
		ID:                   client1.ID,
		FirstName:            client1.FirstName,
		LastName:             client1.LastName,
		Email:                client1.Email,
		Phone:                client1.Phone,
		IdentificationNumber: client1.IdentificationNumber,
		IdentificationType:   client1.IdentificationType,
	}

	client2, err := testQueries.UpdateCustomer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, client2)

	require.Equal(t, client1.ID, client2.ID)
	require.Equal(t, client1.CompanyID, client2.CompanyID)
	require.Equal(t, client1.FirstName, client2.FirstName)
	require.Equal(t, client1.LastName, client2.LastName)
	require.Equal(t, client1.Email, client2.Email)
	require.Equal(t, client1.Phone, client2.Phone)
	require.Equal(t, client1.IdentificationNumber, client2.IdentificationNumber)
	require.Equal(t, client1.IdentificationType, client2.IdentificationType)
	require.WithinDuration(t, client1.CreatedAt, client2.CreatedAt, time.Second)
}

func TestDeleteCustomer(t *testing.T) {
	client1 := createRandomCustomer(t)
	err := testQueries.DeleteCustomer(context.Background(), client1.ID)
	require.NoError(t, err)

	client2, err := testQueries.GetCustomer(context.Background(), client1.ID)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, client2)
}
