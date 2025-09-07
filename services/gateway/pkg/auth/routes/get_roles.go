package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/edorguez/business-manager/services/gateway/pkg/auth/client"
	"github.com/edorguez/business-manager/services/gateway/pkg/auth/contracts"
	"github.com/edorguez/business-manager/services/gateway/pkg/config"
	pb "github.com/edorguez/business-manager/shared/pb/role"
)

func GetRoles(w http.ResponseWriter, r *http.Request, c *config.Config) {
	w.Header().Set("Content-Type", "application/json")

	params := &pb.GetRolesRequest{}

	if err := client.InitRoleServiceClient(c); err != nil {
		json.NewEncoder(w).Encode(&contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	res, err := client.GetRoles(params, r.Context())

	if err != nil {
		fmt.Println("API Gateway :  GetRoles - ERROR")
		json.NewEncoder(w).Encode(err)
		return

	}

	fmt.Println("API Gateway :  GetRoles - SUCCESS")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
