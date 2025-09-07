package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/EdoRguez/business-manager/gateway/pkg/company/client"
	"github.com/EdoRguez/business-manager/gateway/pkg/company/contracts"
	"github.com/EdoRguez/business-manager/gateway/pkg/config"
	"github.com/gorilla/mux"
)

func UpdatePayment(w http.ResponseWriter, r *http.Request, c *config.Config) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("API Gateway :  UpdatePayment")

	// We got our body through context, since we saved it in a middleware
	body := r.Context().Value(contracts.UpdatePaymentRequest{}).(contracts.UpdatePaymentRequest)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		json.NewEncoder(w).Encode(&contracts.Error{
			Status: http.StatusBadRequest,
			Error:  "Unable to convert ID",
		})
	}

	fmt.Println("API Gateway :  UpdatePayment - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	if err := client.InitPaymentServiceClient(c); err != nil {
		json.NewEncoder(w).Encode(&contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	res, errPayment := client.UpdatePayment(int64(id), body, r.Context())

	if errPayment != nil {
		fmt.Println("API Gateway :  UpdatePayment - ERROR")
		json.NewEncoder(w).Encode(errPayment)
		return
	}

	fmt.Println("API Gateway :  UpdatePayment - SUCCESS")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}
