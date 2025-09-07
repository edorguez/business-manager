package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/edorguez/business-manager/services/gateway/pkg/config"
	"github.com/edorguez/business-manager/services/gateway/pkg/product/client"
	"github.com/edorguez/business-manager/services/gateway/pkg/product/contracts"
	"github.com/gorilla/mux"
)

func GetProduct(w http.ResponseWriter, r *http.Request, c *config.Config) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	if err := client.InitProductServiceClient(c); err != nil {
		json.NewEncoder(w).Encode(&contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	res, errProduct := client.GetProduct(vars["id"], r.Context())

	if errProduct != nil {
		fmt.Println("API Gateway :  GetProduct - ERROR")
		json.NewEncoder(w).Encode(errProduct)
		return
	}

	fmt.Println("API Gateway :  GetProduct - SUCCESS")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
