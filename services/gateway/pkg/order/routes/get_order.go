package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/edorguez/business-manager/services/gateway/pkg/config"
	orderClient "github.com/edorguez/business-manager/services/gateway/pkg/order/client"
	"github.com/edorguez/business-manager/shared/types"
	"github.com/gorilla/mux"
)

func GetOrder(w http.ResponseWriter, r *http.Request, c *config.Config) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("API Gateway :  GetOrder")

	// Get order ID from URL path
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		json.NewEncoder(w).Encode(&types.Error{
			Status: http.StatusBadRequest,
			Error:  "Order ID is required",
		})
		return
	}

	if err := orderClient.InitOrderServiceClient(c); err != nil {
		json.NewEncoder(w).Encode(&types.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	res, errResp := orderClient.GetOrder(id, r.Context())
	if errResp != nil {
		fmt.Println("API Gateway :  GetOrder - ERROR")
		json.NewEncoder(w).Encode(errResp)
		return
	}

	fmt.Println("API Gateway :  GetOrder - SUCCESS")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}
