package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/gateway/pkg/config"
	"github.com/EdoRguez/business-manager/gateway/pkg/product/client"
	"github.com/EdoRguez/business-manager/gateway/pkg/product/contracts"
)

func CreateProduct(w http.ResponseWriter, r *http.Request, c *config.Config) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("API Gateway :  CreateProduct")

	// We got our body through context, since we saved it in a middleware
	body := r.Context().Value(contracts.CreateProductRequest{}).(contracts.CreateProductRequest)

	fmt.Println("API Gateway :  CreateProduct - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	if err := client.InitProductServiceClient(c); err != nil {
		json.NewEncoder(w).Encode(&contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	res, err := client.CreateProduct(body, r.Context())
	if err != nil {
		fmt.Println("API Gateway :  CreateProduct - ERROR")
		json.NewEncoder(w).Encode(err)
		return
	}

	fmt.Println("API Gateway :  CreateProduct - SUCCESS")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}
