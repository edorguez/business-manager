package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/gateway/pkg/auth/contracts"
	"github.com/EdoRguez/business-manager/gateway/pkg/config"
	pb "github.com/EdoRguez/business-manager/gateway/pkg/pb/role"
	"google.golang.org/grpc"
)

var roleServiceClient pb.RoleServiceClient

func InitRoleServiceClient(c *config.Config) error {
	fmt.Println("User CLIENT :  InitRoleServiceClient")
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.Auth_Svc_Url, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
		return err
	}

	roleServiceClient = pb.NewRoleServiceClient(cc)
	return nil
}

func GetRole(id int64, c context.Context) (*contracts.GetRoleResponse, *contracts.Error) {
	fmt.Println("Role CLIENT :  GetRole")

	params := &pb.GetRoleRequest{
		Id: id,
	}

	res, err := roleServiceClient.GetRole(c, params)

	if err != nil {
		fmt.Println("Role CLIENT :  GetRole - ERROR")
		fmt.Println(err.Error())

		error := &contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	fmt.Println("Role CLIENT :  GetRole - SUCCESS")

	if res.Status != http.StatusOK {
		error := &contracts.Error{
			Status: res.Status,
			Error:  res.Error,
		}
		return nil, error
	}

	return &contracts.GetRoleResponse{
		Id:   res.Id,
		Name: res.Name,
	}, nil
}

func GetRoles(params *pb.GetRolesRequest, c context.Context) ([]*contracts.GetRoleResponse, *contracts.Error) {
	fmt.Println("Role CLIENT :  GetRoles")

	res, err := roleServiceClient.GetRoles(c, params)

	if err != nil {
		fmt.Println("Role CLIENT :  GetRoles - ERROR")
		fmt.Println(err.Error())

		error := &contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	if res.Status != http.StatusOK {
		error := &contracts.Error{
			Status: res.Status,
			Error:  res.Error,
		}

		return nil, error
	}

	cr := make([]*contracts.GetRoleResponse, 0, len(res.Roles))
	for _, v := range res.Roles {
		cr = append(cr, &contracts.GetRoleResponse{
			Id:   v.Id,
			Name: v.Name,
		})
	}

	fmt.Println("Role CLIENT :  GetRoles - SUCCESS")
	return cr, nil
}
