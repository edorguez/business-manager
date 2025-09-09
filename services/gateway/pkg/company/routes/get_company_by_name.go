package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/edorguez/business-manager/services/gateway/pkg/company/client"
	"github.com/edorguez/business-manager/services/gateway/pkg/config"
	"github.com/edorguez/business-manager/shared/types"
	"github.com/gorilla/mux"
)

func GetCompanyByName(w http.ResponseWriter, r *http.Request, c *config.Config) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	name := vars["name"]
	if name == "" {
		json.NewEncoder(w).Encode(&types.Error{
			Status: http.StatusBadRequest,
			Error:  "Name param must be provided",
		})
		return
	}

	if err := client.InitCompanyServiceClient(c); err != nil {
		json.NewEncoder(w).Encode(&types.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	res, errCompany := client.GetCompanyByName(strings.ToLower(name), r.Context())

	if errCompany != nil {
		fmt.Println("API Gateway :  GetCompanyByName - ERROR")
		json.NewEncoder(w).Encode(errCompany)
		return
	}

	fmt.Println("API Gateway :  GetCompanyByName - SUCCESS")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
