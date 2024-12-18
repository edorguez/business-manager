package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/gateway/pkg/config"
	pb "github.com/EdoRguez/business-manager/gateway/pkg/pb/product"
	"github.com/EdoRguez/business-manager/gateway/pkg/product/client"
	"github.com/EdoRguez/business-manager/gateway/pkg/product/contracts"
	"github.com/EdoRguez/business-manager/gateway/pkg/util/query_params"
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
		json.NewEncoder(w).Encode(&contracts.Error{
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
