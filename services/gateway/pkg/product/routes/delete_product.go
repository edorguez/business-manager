package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/edorguez/business-manager/services/gateway/pkg/config"
	"github.com/edorguez/business-manager/services/gateway/pkg/product/client"
	"github.com/edorguez/business-manager/shared/types"
	"github.com/gorilla/mux"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request, c *config.Config) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	if err := client.InitProductServiceClient(c); err != nil {
		json.NewEncoder(w).Encode(&types.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	res, errDelete := client.DeleteProduct(vars["id"], r.Context())

	if errDelete != nil {
		fmt.Println("API Gateway :  DeleteProduct - ERROR")
		json.NewEncoder(w).Encode(errDelete)
		return
	}

	fmt.Println("API Gateway :  DeleteProduct - SUCCESS")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}
