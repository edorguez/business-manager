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

func UpdateCompany(w http.ResponseWriter, r *http.Request, c *config.Config) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("API Gateway :  UpdateCompany")

	// We got our body through context, since we saved it in a middleware
	body := r.Context().Value(contracts.UpdateCompanyRequest{}).(contracts.UpdateCompanyRequest)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		json.NewEncoder(w).Encode(&contracts.Error{
			Status: http.StatusBadRequest,
			Error:  "Unable to convert ID",
		})
	}

	fmt.Println("API Gateway :  UpdateCompany - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	if err := client.InitCompanyServiceClient(c); err != nil {
		json.NewEncoder(w).Encode(&contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	res, errCompany := client.UpdateCompany(int64(id), body, r.Context())

	if errCompany != nil {
		fmt.Println("API Gateway :  UpdateCompany - ERROR")
		json.NewEncoder(w).Encode(errCompany)
		return
	}

	fmt.Println("API Gateway :  UpdateCompany - SUCCESS")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}
