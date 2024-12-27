package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/gateway/pkg/auth/client"
	"github.com/EdoRguez/business-manager/gateway/pkg/auth/contracts"
	"github.com/EdoRguez/business-manager/gateway/pkg/config"
)

func Login(w http.ResponseWriter, r *http.Request, c *config.Config) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("API Gateway :  Login")

	// We got our body through context, since we saved it in a middleware
	body := r.Context().Value(contracts.LoginRequest{}).(contracts.LoginRequest)

	fmt.Println("API Gateway :  Login - Body")
	// fmt.Println(body)
	fmt.Println("-----------------")

	if err := client.InitAuthServiceClient(c); err != nil {
		json.NewEncoder(w).Encode(&contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	res, err := client.Login(body, r.Context())

	if err != nil {
		fmt.Println("API Gateway :  Login - ERROR")
		json.NewEncoder(w).Encode(err)
		return
	}

	fmt.Println("API Gateway :  Login - SUCCESS")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}
