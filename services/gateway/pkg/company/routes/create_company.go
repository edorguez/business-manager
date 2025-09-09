package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/edorguez/business-manager/services/gateway/pkg/company/client"
	"github.com/edorguez/business-manager/services/gateway/pkg/company/contracts"
	"github.com/edorguez/business-manager/services/gateway/pkg/config"
	"github.com/edorguez/business-manager/shared/types"
)

func CreateCompany(w http.ResponseWriter, r *http.Request, c *config.Config) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("API Gateway :  CreateCompany")

	// We got our body through context, since we saved it in a middleware
	body := r.Context().Value(contracts.CreateCompanyRequest{}).(contracts.CreateCompanyRequest)

	fmt.Println("API Gateway :  CreateCompany - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	if err := client.InitCompanyServiceClient(c); err != nil {
		json.NewEncoder(w).Encode(&types.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	res, err := client.CreateCompany(body, r.Context())

	if err != nil {
		fmt.Println("API Gateway :  CreateCompany - ERROR")
		json.NewEncoder(w).Encode(err)
		return
	}

	fmt.Println("API Gateway :  CreateCompany - SUCCESS")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}
