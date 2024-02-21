package services

import (
	"context"
	"fmt"
	"net/http"

	db "github.com/EdoRguez/business-manager/client-svc/pkg/db/sqlc"
	client "github.com/EdoRguez/business-manager/client-svc/pkg/pb"
	repo "github.com/EdoRguez/business-manager/client-svc/pkg/repository"
	util "github.com/EdoRguez/business-manager/client-svc/pkg/util"
)

type ClientService struct {
	Repo *repo.ClientRepo
	client.UnimplementedClientServiceServer
}

func (s *ClientService) CreateClient(ctx context.Context, req *client.CreateClientRequest) (*client.CreateClientResponse, error) {
	fmt.Println("Client Service :  CreateClient")
	fmt.Println("Client Service :  Create Client - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	createClientParams := db.CreateClientParams{
		CompanyID:            req.CompanyId,
		FirstName:            req.FirstName,
		LastName:             util.NewSqlNullString(req.LastName),
		Email:                util.NewSqlNullString(req.Email),
		Phone:                util.NewSqlNullString(req.Phone),
		IdentificationNumber: req.IdentificationNumber,
		IdentificationType:   req.IdentificationType,
	}

	c, err := s.Repo.CreateClient(ctx, createClientParams)
	if err != nil {
		fmt.Println("API Gateway :  Create Client - ERROR")
		fmt.Println(err.Error())
		return &client.CreateClientResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	fmt.Println("API Gateway :  Create Client - SUCCESS")
	return &client.CreateClientResponse{
		Status: http.StatusCreated,
		Id:     c.ID,
	}, nil
}
