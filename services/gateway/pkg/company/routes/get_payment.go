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

func GetPayment(w http.ResponseWriter, r *http.Request, c *config.Config) {
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

	if err := client.InitPaymentServiceClient(c); err != nil {
		json.NewEncoder(w).Encode(&contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	res, errPayment := client.GetPayment(int64(id), r.Context())

	if errPayment != nil {
		fmt.Println("API Gateway :  GetPayment - ERROR")
		json.NewEncoder(w).Encode(errPayment)
		return
	}

	fmt.Println("API Gateway :  GetPayment - SUCCESS")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
