package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/edorguez/business-manager/services/gateway/pkg/config"
	"github.com/edorguez/business-manager/services/gateway/pkg/customer/client"
	"github.com/edorguez/business-manager/services/gateway/pkg/customer/contracts"
	"github.com/gorilla/mux"
)

func GetCustomer(w http.ResponseWriter, r *http.Request, c *config.Config) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		json.NewEncoder(w).Encode(&contracts.Error{
			Status: http.StatusBadRequest,
			Error:  "Unable to convert ID",
		})
		return
	}

	if err := client.InitCustomerServiceClient(c); err != nil {
		json.NewEncoder(w).Encode(&contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	res, errCustomer := client.GetCustomer(int64(id), r.Context())

	if errCustomer != nil {
		fmt.Println("API Gateway :  GetCustomer - ERROR")
		json.NewEncoder(w).Encode(errCustomer)
		return
	}

	fmt.Println("API Gateway :  GetCustomer - SUCCESS")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
