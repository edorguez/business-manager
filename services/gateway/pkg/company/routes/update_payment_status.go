package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/edorguez/business-manager/services/gateway/pkg/company/client"
	"github.com/edorguez/business-manager/services/gateway/pkg/company/contracts"
	"github.com/edorguez/business-manager/services/gateway/pkg/config"
	"github.com/gorilla/mux"
)

func UpdatePaymentStatus(w http.ResponseWriter, r *http.Request, c *config.Config) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("API Gateway :  UpdatePaymentStatus")

	// We got our body through context, since we saved it in a middleware
	// body := r.Context().Value(contracts.UpdatePaymentStatusRequest{}).(contracts.UpdatePaymentStatusRequest)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		json.NewEncoder(w).Encode(&contracts.Error{
			Status: http.StatusBadRequest,
			Error:  "Unable to convert ID",
		})
	}

	fmt.Println("API Gateway :  UpdatePaymentStatus - Body")
	fmt.Println(r.Body)
	fmt.Println("-----------------")

	var body contracts.UpdatePaymentStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := client.InitPaymentServiceClient(c); err != nil {
		json.NewEncoder(w).Encode(&contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	res, errPayment := client.UpdatePaymentStatus(int64(id), body, r.Context())

	if errPayment != nil {
		fmt.Println("API Gateway :  UpdatePaymentStatus - ERROR")
		json.NewEncoder(w).Encode(errPayment)
		return
	}

	fmt.Println("API Gateway :  UpdatePaymentStatus - SUCCESS")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}
