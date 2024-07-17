package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/EdoRguez/business-manager/gateway/pkg/auth/client"
	"github.com/EdoRguez/business-manager/gateway/pkg/auth/contracts"
	"github.com/EdoRguez/business-manager/gateway/pkg/config"
	"github.com/gorilla/mux"
)

func UpdateUser(w http.ResponseWriter, r *http.Request, c *config.Config) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("API Gateway :  UpdateUser")

	// We got our body through context, since we saved it in a middleware
	body := r.Context().Value(contracts.UpdateUserRequest{}).(contracts.UpdateUserRequest)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		json.NewEncoder(w).Encode(&contracts.Error{
			Status: http.StatusBadRequest,
			Error:  "Unable to convert ID",
		})
	}

	fmt.Println("API Gateway :  UpdateUser - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	if err := client.InitUserServiceClient(c); err != nil {
		json.NewEncoder(w).Encode(&contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	res, errUpdate := client.UpdateUser(int64(id), body, r.Context())

	if errUpdate != nil {
		fmt.Println("API Gateway :  UpdateUser - ERROR")
		json.NewEncoder(w).Encode(errUpdate)
		return
	}

	fmt.Println("API Gateway :  UpdateUser - SUCCESS")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}
