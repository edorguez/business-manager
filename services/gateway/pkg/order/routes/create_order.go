package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	companyClient "github.com/edorguez/business-manager/services/gateway/pkg/company/client"
	"github.com/edorguez/business-manager/services/gateway/pkg/config"
	orderClient "github.com/edorguez/business-manager/services/gateway/pkg/order/client"
	"github.com/edorguez/business-manager/services/gateway/pkg/order/contracts"
	"github.com/edorguez/business-manager/shared/types"
)

func CreateOrder(w http.ResponseWriter, r *http.Request, c *config.Config) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("API Gateway :  CreateOrder")

	// We got our body through context, since we saved it in a middleware
	body := r.Context().Value("keyOrderCreate").(contracts.CreateOrderRequest)

	fmt.Println("API Gateway :  CreateOrder - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	if err := orderClient.InitOrderServiceClient(c); err != nil {
		json.NewEncoder(w).Encode(&types.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	if err := companyClient.InitCompanyServiceClient(c); err != nil {
		json.NewEncoder(w).Encode(&types.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	_, errCompany := companyClient.GetCompany(int64(body.CompanyId), r.Context())
	if errCompany != nil {
		fmt.Println("API Gateway :  CreateOrder - ERROR")
		errCompany.Error = "Company not found"
		json.NewEncoder(w).Encode(errCompany)
		return
	}

	res, err := orderClient.CreateOrder(body, r.Context())
	if err != nil {
		fmt.Println("API Gateway :  CreateOrder - ERROR")
		json.NewEncoder(w).Encode(err)
		return
	}

	fmt.Println("API Gateway :  CreateOrder - SUCCESS")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}
