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
	fmt.Println("Client Service :  CreateClient - Req")
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
		fmt.Println("API Gateway :  CreateClient - ERROR")
		fmt.Println(err.Error())
		return &client.CreateClientResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	fmt.Println("API Gateway :  CreateClient - SUCCESS")
	return &client.CreateClientResponse{
		Status: http.StatusCreated,
		Id:     c.ID,
	}, nil
}

func (s *ClientService) GetClient(ctx context.Context, req *client.GetClientRequest) (*client.GetClientResponse, error) {
	fmt.Println("Client Service :  GetClient")
	fmt.Println("Client Service :  GetClient - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	c, err := s.Repo.GetClient(ctx, req.Id)
	if err != nil {
		fmt.Println("API Gateway :  GetClient - ERROR")
		fmt.Println(err.Error())
		return &client.GetClientResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	fmt.Println("API Gateway :  GetClient - SUCCESS")
	return &client.GetClientResponse{
		Id:                   c.ID,
		CompanyId:            c.CompanyID,
		FirstName:            c.FirstName,
		LastName:             c.LastName.String,
		Email:                c.Email.String,
		Phone:                c.Phone.String,
		IdentificationNumber: c.IdentificationNumber,
		IdentificationType:   c.IdentificationType,
		Status:               http.StatusOK,
	}, nil
}

func (s *ClientService) GetClients(ctx context.Context, req *client.GetClientsRequest) (*client.GetClientsResponse, error) {
	fmt.Println("Client Service :  GetClients")
	fmt.Println("Client Service :  GetClients - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	params := db.GetClientsParams{
		CompanyID: req.CompanyId,
		Limit:     req.Limit,
		Offset:    req.Offset,
	}

	c, err := s.Repo.GetClients(ctx, params)
	if err != nil {
		fmt.Println("API Gateway :  GetClients - ERROR")
		fmt.Println(err.Error())
		return &client.GetClientsResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	var clients []*client.GetClientResponse
	for _, v := range c {
		clients = append(clients, &client.GetClientResponse{
			Id:                   v.ID,
			CompanyId:            v.CompanyID,
			FirstName:            v.FirstName,
			LastName:             v.LastName.String,
			Email:                v.Email.String,
			Phone:                v.Phone.String,
			IdentificationNumber: v.IdentificationNumber,
			IdentificationType:   v.IdentificationType,
			Status:               http.StatusOK,
		})
	}

	fmt.Println("API Gateway :  GetClients - SUCCESS")
	return &client.GetClientsResponse{
		Clients: clients,
		Status:  http.StatusOK,
	}, nil
}

func (s *ClientService) UpdateClient(ctx context.Context, req *client.UpdateClientRequest) (*client.UpdateClientResponse, error) {
	fmt.Println("Client Service :  UpdateClient")
	fmt.Println("Client Service :  UpdateClient - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	params := db.UpdateClientParams{
		ID:                   req.Id,
		FirstName:            req.FirstName,
		LastName:             util.NewSqlNullString(req.LastName),
		Email:                util.NewSqlNullString(req.Email),
		Phone:                util.NewSqlNullString(req.Phone),
		IdentificationNumber: req.IdentificationNumber,
		IdentificationType:   req.IdentificationType,
	}

	_, err := s.Repo.UpdateClient(ctx, params)
	if err != nil {
		fmt.Println("API Gateway :  UpdateClient - ERROR")
		fmt.Println(err.Error())
		return &client.UpdateClientResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	fmt.Println("API Gateway :  UpdateClient - SUCCESS")
	return &client.UpdateClientResponse{
		Status: http.StatusNoContent,
	}, nil
}

func (s *ClientService) DeleteClient(ctx context.Context, req *client.DeleteClientRequest) (*client.DeleteClientResponse, error) {
	fmt.Println("Client Service :  DeleteClient")
	fmt.Println("Client Service :  DeleteClient - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	err := s.Repo.DeleteClient(ctx, req.Id)
	if err != nil {
		fmt.Println("API Gateway :  DeleteClient - ERROR")
		fmt.Println(err.Error())
		return &client.DeleteClientResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	fmt.Println("API Gateway :  DeleteClient - SUCCESS")
	return &client.DeleteClientResponse{
		Status: http.StatusNoContent,
	}, nil
}
