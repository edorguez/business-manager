package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/EdoRguez/business-manager/client-svc/pkg/util"
	"github.com/stretchr/testify/require"
)

func createRandomClient(t *testing.T) ClientClient {
	arg := CreateClientParams{
		CompanyID:            1,
		FirstName:            util.RandomName(),
		LastName:             util.NewSqlNullString(util.RandomName()),
		Email:                util.NewSqlNullString(util.RandomEmail()),
		Phone:                util.NewSqlNullString(util.RandomPhoneNumber()),
		IdentificationNumber: util.RandomIndentificationNumber(),
		IdentificationType:   util.RandomIdentificationType(),
	}

	client, err := testQueries.CreateClient(context.Background(), arg)
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

func TestCreateClient(t *testing.T) {
	createRandomClient(t)
}

func TestGetClient(t *testing.T) {
	client1 := createRandomClient(t)
	client2, err := testQueries.GetClient(context.Background(), client1.ID)
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

func TestGetClients(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomClient(t)
	}

	arg := GetClientsParams{
		CompanyID: 0,
		Limit:     5,
		Offset:    5,
	}

	clients, err := testQueries.GetClients(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, clients, 5)

	for _, client := range clients {
		require.NotEmpty(t, client)
	}
}

func TestUpdateClient(t *testing.T) {
	client1 := createRandomClient(t)

	arg := UpdateClientParams{
		ID:                   client1.ID,
		FirstName:            client1.FirstName,
		LastName:             client1.LastName,
		Email:                client1.Email,
		Phone:                client1.Phone,
		IdentificationNumber: client1.IdentificationNumber,
		IdentificationType:   client1.IdentificationType,
	}

	client2, err := testQueries.UpdateClient(context.Background(), arg)
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

func TestDeleteClient(t *testing.T) {
	client1 := createRandomClient(t)
	err := testQueries.DeleteClient(context.Background(), client1.ID)
	require.NoError(t, err)

	client2, err := testQueries.GetClient(context.Background(), client1.ID)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, client2)
}
