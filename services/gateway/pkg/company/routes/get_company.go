package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/edorguez/business-manager/services/gateway/pkg/company/client"
	"github.com/edorguez/business-manager/services/gateway/pkg/config"
	"github.com/edorguez/business-manager/shared/types"
	"github.com/gorilla/mux"
)

func GetCompany(w http.ResponseWriter, r *http.Request, c *config.Config) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		json.NewEncoder(w).Encode(&types.Error{
			Status: http.StatusBadRequest,
			Error:  "Unable to convert ID",
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

	res, errCompany := client.GetCompany(int64(id), r.Context())

	if errCompany != nil {
		fmt.Println("API Gateway :  GetCompany - ERROR")
		json.NewEncoder(w).Encode(errCompany)
		return
	}

	fmt.Println("API Gateway :  GetCompany - SUCCESS")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
