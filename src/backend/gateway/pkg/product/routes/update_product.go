package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/gateway/pkg/config"
	"github.com/EdoRguez/business-manager/gateway/pkg/product/client"
	"github.com/EdoRguez/business-manager/gateway/pkg/product/contracts"
	"github.com/gorilla/mux"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request, c *config.Config) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("API Gateway :  UpdateProduct")

	// We got our body through context, since we saved it in a middleware
	body := r.Context().Value("keyProductUpdate").(contracts.UpdateProductRequest)

	vars := mux.Vars(r)

	fmt.Println("API Gateway :  UpdateProduct - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	if err := client.InitProductServiceClient(c); err != nil {
		json.NewEncoder(w).Encode(&contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	res, errUpdate := client.UpdateProduct(vars["id"], body, r.Context())

	if errUpdate != nil {
		fmt.Println("API Gateway :  UpdateProduct - ERROR")
		json.NewEncoder(w).Encode(errUpdate)
		return
	}

	fmt.Println("API Gateway :  UpdateProduct - SUCCESS")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}
