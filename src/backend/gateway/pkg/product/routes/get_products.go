package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/EdoRguez/business-manager/gateway/pkg/config"
	pb "github.com/EdoRguez/business-manager/gateway/pkg/pb/product"
	"github.com/EdoRguez/business-manager/gateway/pkg/product/client"
	"github.com/EdoRguez/business-manager/gateway/pkg/product/contracts"
	"github.com/EdoRguez/business-manager/gateway/pkg/util/query_params"
)

func GetProducts(w http.ResponseWriter, r *http.Request, c *config.Config) {
	w.Header().Set("Content-Type", "application/json")
	companyId := query_params.GetId("companyId", r)
	name := r.URL.Query().Get("name")
	sku := r.URL.Query().Get("sku")
	productStatusStr := r.URL.Query().Get("productStatus")
	var productStatus *uint32
	if productStatusStr != "" {
		status, err := strconv.ParseUint(productStatusStr, 10, 32)
		if err == nil {
			status32 := uint32(status)
			productStatus = &status32
		}
	}
	limit, offset := query_params.GetFilter(r)

	params := &pb.GetProductsRequest{
		CompanyId:     companyId,
		Name:          &name,
		Sku:           &sku,
		ProductStatus: productStatus,
		Limit:         limit,
		Offset:        offset,
	}

	if err := client.InitProductServiceClient(c); err != nil {
		json.NewEncoder(w).Encode(&contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	res, err := client.GetProducts(params, r.Context())

	if err != nil {
		fmt.Println("API Gateway :  GetProducts - ERROR")
		json.NewEncoder(w).Encode(err)
		return

	}

	fmt.Println("API Gateway :  GetProducts - SUCCESS")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
