package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/gateway/pkg/auth/client"
	"github.com/EdoRguez/business-manager/gateway/pkg/auth/contracts"
	"github.com/EdoRguez/business-manager/gateway/pkg/config"
)

func Register(w http.ResponseWriter, r *http.Request, c *config.Config) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("API Gateway :  Register")

	// We got our body through context, since we saved it in a middleware
	body := r.Context().Value(contracts.CreateUserRequest{}).(contracts.CreateUserRequest)

	fmt.Println("API Gateway :  Register - Body")
	// fmt.Println(body)
	fmt.Println("-----------------")

	if err := client.InitAuthServiceClient(c); err != nil {
		json.NewEncoder(w).Encode(&contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	res, err := client.Register(body, r.Context())

	if err != nil {
		fmt.Println("API Gateway :  Register - ERROR")
		json.NewEncoder(w).Encode(err)
		return
	}

	fmt.Println("API Gateway :  Register - SUCCESS")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}
