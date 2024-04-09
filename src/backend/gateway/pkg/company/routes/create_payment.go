package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/gateway/pkg/company/client"
	"github.com/EdoRguez/business-manager/gateway/pkg/company/contracts"
	"github.com/EdoRguez/business-manager/gateway/pkg/config"
)

func CreatePayment(w http.ResponseWriter, r *http.Request, c *config.Config) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("API Gateway :  CreatePayment")

	// We got our body through context, since we saved it in a middleware
	body := r.Context().Value(contracts.CreatePaymentRequest{}).(contracts.CreatePaymentRequest)

	fmt.Println("API Gateway :  CreatePayment - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	if err := client.InitPaymentServiceClient(c); err != nil {
		json.NewEncoder(w).Encode(&contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	res, err := client.CreatePayment(body, r.Context())

	if err != nil {
		fmt.Println("API Gateway :  CreatePayment - ERROR")
		json.NewEncoder(w).Encode(err)
		return
	}

	fmt.Println("API Gateway :  CreatePayment - SUCCESS")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}
