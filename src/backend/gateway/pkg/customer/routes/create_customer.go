package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/gateway/pkg/config"
	"github.com/EdoRguez/business-manager/gateway/pkg/customer/client"
	"github.com/EdoRguez/business-manager/gateway/pkg/customer/contracts"
)

func CreateCustomer(w http.ResponseWriter, r *http.Request, c *config.Config) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("API Gateway :  CreateCustomer")

	// We got our body through context, since we saved it in a middleware
	body := r.Context().Value(contracts.CreateCustomerRequest{}).(contracts.CreateCustomerRequest)

	fmt.Println("API Gateway :  CreateCustomer - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	if err := client.InitCustomerServiceClient(c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	res, err := client.CreateCustomer(body, r.Context())

	if err != nil {
		fmt.Println("API Gateway :  CreateCustomer - ERROR")
		http.Error(w, err.Error, int(err.Status))
		json.NewEncoder(w).Encode(err)
		return
	}

	fmt.Println("API Gateway :  CreateCustomer - SUCCESS")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}
