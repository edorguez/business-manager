package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/edorguez/business-manager/services/gateway/pkg/config"
	"github.com/edorguez/business-manager/services/gateway/pkg/product/client"
	pb "github.com/edorguez/business-manager/shared/pb/product"
	"github.com/edorguez/business-manager/shared/types"
	"github.com/edorguez/business-manager/shared/util/query_params"
)

func GetLatestProducts(w http.ResponseWriter, r *http.Request, c *config.Config) {
	w.Header().Set("Content-Type", "application/json")
	companyId := query_params.GetId("companyId", r)
	limit, _ := query_params.GetFilter(r)

	params := &pb.GetLatestProductsRequest{
		CompanyId: companyId,
		Limit:     limit,
	}

	if err := client.InitProductServiceClient(c); err != nil {
		json.NewEncoder(w).Encode(&types.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	res, err := client.GetLatestProducts(params, r.Context())

	if err != nil {
		fmt.Println("API Gateway :  GetLatestProducts - ERROR")
		json.NewEncoder(w).Encode(err)
		return

	}

	fmt.Println("API Gateway :  GetLatestProducts - SUCCESS")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
