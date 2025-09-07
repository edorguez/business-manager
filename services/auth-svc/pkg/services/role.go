package services

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	repo "github.com/edorguez/business-manager/services/auth-svc/pkg/repository"
	pb "github.com/edorguez/business-manager/shared/pb/role"
	"github.com/edorguez/business-manager/shared/util/type_converter"
)

type RoleService struct {
	Repo *repo.RoleRepo
	pb.UnimplementedRoleServiceServer
}

func (s *RoleService) GetRole(ctx context.Context, req *pb.GetRoleRequest) (*pb.GetRoleResponse, error) {
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

		return &pb.GetRoleResponse{
			Status: int64(resErrorStatus),
			Error:  resErrorMessage,
		}, nil
	}

	fmt.Println("Auth Service :  GetRole - SUCCESS")
	return &pb.GetRoleResponse{
		Id:          c.ID,
		Name:        c.Name,
		Description: type_converter.NewString(c.Description),
		Status:      http.StatusOK,
	}, nil
}

func (s *RoleService) GetRoles(ctx context.Context, req *pb.GetRolesRequest) (*pb.GetRolesResponse, error) {
	fmt.Println("Auth Service :  GetRoles")
	fmt.Println("Auth Service :  GetRoles - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	c, err := s.Repo.GetRoles(ctx)
	if err != nil {
		fmt.Println("Auth Service :  GetRoles - ERROR")
		fmt.Println(err.Error())
		return &pb.GetRolesResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	roles := make([]*pb.GetRoleResponse, 0, len(c))
	for _, v := range c {
		roles = append(roles, &pb.GetRoleResponse{
			Id:          v.ID,
			Name:        v.Name,
			Description: type_converter.NewString(v.Description),
			Status:      http.StatusOK,
		})
	}

	fmt.Println("Auth Service :  GetRoles - SUCCESS")
	return &pb.GetRolesResponse{
		Roles:  roles,
		Status: http.StatusOK,
	}, nil
}
