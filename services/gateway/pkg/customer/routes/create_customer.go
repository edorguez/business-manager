package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	companyClient "github.com/edorguez/business-manager/services/gateway/pkg/company/client"
	"github.com/edorguez/business-manager/services/gateway/pkg/config"
	customerClient "github.com/edorguez/business-manager/services/gateway/pkg/customer/client"
	"github.com/edorguez/business-manager/services/gateway/pkg/customer/contracts"
)

func CreateCustomer(w http.ResponseWriter, r *http.Request, c *config.Config) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("API Gateway :  CreateCustomer")

	// We got our body through context, since we saved it in a middleware
	body := r.Context().Value(contracts.CreateCustomerRequest{}).(contracts.CreateCustomerRequest)

	fmt.Println("API Gateway :  CreateCustomer - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	if err := customerClient.InitCustomerServiceClient(c); err != nil {
		json.NewEncoder(w).Encode(&contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	if err := companyClient.InitCompanyServiceClient(c); err != nil {
		json.NewEncoder(w).Encode(&contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	_, errCompany := companyClient.GetCompany(int64(body.CompanyId), r.Context())
	if errCompany != nil {
		fmt.Println("API Gateway :  CreateCustomer - ERROR")
		errCompany.Error = "Company not found"
		json.NewEncoder(w).Encode(errCompany)
		return
	}

	res, err := customerClient.CreateCustomer(body, r.Context())
	if err != nil {
		fmt.Println("API Gateway :  CreateCustomer - ERROR")
		json.NewEncoder(w).Encode(err)
		return
	}

	fmt.Println("API Gateway :  CreateCustomer - SUCCESS")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}
