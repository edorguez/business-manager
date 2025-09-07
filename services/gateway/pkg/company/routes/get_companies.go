package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/edorguez/business-manager/services/gateway/pkg/company/client"
	"github.com/edorguez/business-manager/services/gateway/pkg/company/contracts"
	"github.com/edorguez/business-manager/services/gateway/pkg/config"
	pb "github.com/edorguez/business-manager/shared/pb/company"
	"github.com/edorguez/business-manager/shared/util/query_params"
)

func GetCompanies(w http.ResponseWriter, r *http.Request, c *config.Config) {
	w.Header().Set("Content-Type", "application/json")
	limit, offset := query_params.GetFilter(r)

	params := &pb.GetCompaniesRequest{
		Limit:  limit,
		Offset: offset,
	}

	if err := client.InitCompanyServiceClient(c); err != nil {
		json.NewEncoder(w).Encode(&contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	res, err := client.GetCompanies(params, r.Context())

	if err != nil {
		fmt.Println("API Gateway :  GetCompanies - ERROR")
		json.NewEncoder(w).Encode(err)
		return

	}

	fmt.Println("API Gateway :  GetCompanies - SUCCESS")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
