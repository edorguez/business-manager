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

func DeleteCustomer(w http.ResponseWriter, r *http.Request, c *config.Config) {
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

	res, errDelete := client.DeleteCustomer(int64(id), r.Context())

	if errDelete != nil {
		fmt.Println("API Gateway :  DeleteCustomer - ERROR")
		json.NewEncoder(w).Encode(errDelete)
		return
	}

	fmt.Println("API Gateway :  DeleteCustomer - SUCCESS")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}
