package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/edorguez/business-manager/services/gateway/pkg/auth/client"
	"github.com/edorguez/business-manager/services/gateway/pkg/config"
	pb "github.com/edorguez/business-manager/shared/pb/user"
	"github.com/edorguez/business-manager/shared/types"
	"github.com/edorguez/business-manager/shared/util/query_params"
)

func GetUsers(w http.ResponseWriter, r *http.Request, c *config.Config) {
	w.Header().Set("Content-Type", "application/json")
	companyId := query_params.GetId("companyId", r)
	limit, offset := query_params.GetFilter(r)

	params := &pb.GetUsersRequest{
		CompanyId: companyId,
		Limit:     limit,
		Offset:    offset,
	}

	if err := client.InitUserServiceClient(c); err != nil {
		json.NewEncoder(w).Encode(&types.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	res, err := client.GetUsers(params, r.Context())

	if err != nil {
		fmt.Println("API Gateway :  GetUsers - ERROR")
		json.NewEncoder(w).Encode(err)
		return

	}

	fmt.Println("API Gateway :  GetUsers - SUCCESS")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
