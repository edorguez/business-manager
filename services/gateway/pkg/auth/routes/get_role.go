package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/edorguez/business-manager/services/gateway/pkg/auth/client"
	"github.com/edorguez/business-manager/services/gateway/pkg/auth/contracts"
	"github.com/edorguez/business-manager/services/gateway/pkg/config"
	"github.com/gorilla/mux"
)

func GetRole(w http.ResponseWriter, r *http.Request, c *config.Config) {
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

	if err := client.InitRoleServiceClient(c); err != nil {
		json.NewEncoder(w).Encode(&contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	res, errRole := client.GetRole(int64(id), r.Context())

	if errRole != nil {
		fmt.Println("API Gateway :  GetRole - ERROR")
		json.NewEncoder(w).Encode(errRole)
		return
	}

	fmt.Println("API Gateway :  GetRole - SUCCESS")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
