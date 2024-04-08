package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/EdoRguez/business-manager/gateway/pkg/config"
	"github.com/EdoRguez/business-manager/gateway/pkg/customer/client"
	"github.com/EdoRguez/business-manager/gateway/pkg/customer/contracts"
	"github.com/gorilla/mux"
)

func UpdateCustomer(w http.ResponseWriter, r *http.Request, c *config.Config) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("API Gateway :  UpdateCustomer")

	// We got our body through context, since we saved it in a middleware
	body := r.Context().Value(contracts.UpdateCustomerRequest{}).(contracts.UpdateCustomerRequest)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		json.NewEncoder(w).Encode(&contracts.Error{
			Status: http.StatusBadRequest,
			Error:  "Unable to convert ID",
		})
	}

	fmt.Println("API Gateway :  UpdateCustomer - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	if err := client.InitCustomerServiceClient(c); err != nil {
		json.NewEncoder(w).Encode(&contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	res, errUpdate := client.UpdateCustomer(int64(id), body, r.Context())

	if errUpdate != nil {
		fmt.Println("API Gateway :  UpdateCustomer - ERROR")
		json.NewEncoder(w).Encode(errUpdate)
		return
	}

	fmt.Println("API Gateway :  UpdateCustomer - SUCCESS")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}
