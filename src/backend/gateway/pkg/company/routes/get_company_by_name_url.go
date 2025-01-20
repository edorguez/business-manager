package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/EdoRguez/business-manager/gateway/pkg/company/client"
	"github.com/EdoRguez/business-manager/gateway/pkg/company/contracts"
	"github.com/EdoRguez/business-manager/gateway/pkg/config"
	"github.com/gorilla/mux"
)

func GetCompanyByNameUrl(w http.ResponseWriter, r *http.Request, c *config.Config) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	nameUrl := vars["nameUrl"]
	if nameUrl == "" {
		json.NewEncoder(w).Encode(&contracts.Error{
			Status: http.StatusBadRequest,
			Error:  "Name URL param must be provided",
		})
		return
	}

	if err := client.InitCompanyServiceClient(c); err != nil {
		json.NewEncoder(w).Encode(&contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	res, errCompany := client.GetCompanyByNameUrl(strings.ToLower(nameUrl), r.Context())

	if errCompany != nil {
		fmt.Println("API Gateway :  GetCompanyByNameUrl - ERROR")
		json.NewEncoder(w).Encode(errCompany)
		return
	}

	fmt.Println("API Gateway :  GetCompanyByNameUrl - SUCCESS")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
