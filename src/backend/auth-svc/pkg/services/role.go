package services

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	auth "github.com/EdoRguez/business-manager/auth-svc/pkg/pb"
	repo "github.com/EdoRguez/business-manager/auth-svc/pkg/repository"
	"github.com/EdoRguez/business-manager/auth-svc/pkg/util/type_converter"
)

type RoleService struct {
	Repo *repo.RoleRepo
	auth.UnimplementedRoleServiceServer
}

func (s *RoleService) GetRole(ctx context.Context, req *auth.GetRoleRequest) (*auth.GetRoleResponse, error) {
	fmt.Println("Auth Service :  GetRole")
	fmt.Println("Auth Service :  GetRole - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	c, err := s.Repo.GetRole(ctx, req.Id)
	if err != nil {
		fmt.Println("Auth Service :  GetRole - ERROR")
		fmt.Println(err.Error())

		resErrorStatus := http.StatusConflict
		resErrorMessage := err.Error()

		if err == sql.ErrNoRows {
			resErrorStatus = http.StatusNotFound
			resErrorMessage = "Record not found"
		}

		return &auth.GetRoleResponse{
			Status: int64(resErrorStatus),
			Error:  resErrorMessage,
		}, nil
	}

	fmt.Println("Auth Service :  GetRole - SUCCESS")
	return &auth.GetRoleResponse{
		Id:          c.ID,
		Name:        c.Name,
		Description: type_converter.NewString(c.Description),
		Status:      http.StatusOK,
	}, nil
}

func (s *RoleService) GetRoles(ctx context.Context, req *auth.GetRolesRequest) (*auth.GetRolesResponse, error) {
	fmt.Println("Auth Service :  GetRoles")
	fmt.Println("Auth Service :  GetRoles - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	c, err := s.Repo.GetRoles(ctx)
	if err != nil {
		fmt.Println("Auth Service :  GetRoles - ERROR")
		fmt.Println(err.Error())
		return &auth.GetRolesResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	roles := make([]*auth.GetRoleResponse, 0, len(c))
	for _, v := range c {
		roles = append(roles, &auth.GetRoleResponse{
			Id:          v.ID,
			Name:        v.Name,
			Description: type_converter.NewString(v.Description),
			Status:      http.StatusOK,
		})
	}

	fmt.Println("Auth Service :  GetRoles - SUCCESS")
	return &auth.GetRolesResponse{
		Roles:  roles,
		Status: http.StatusOK,
	}, nil
}
