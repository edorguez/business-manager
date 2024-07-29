package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/gateway/pkg/config"
	"github.com/EdoRguez/business-manager/gateway/pkg/customer/client"
	"github.com/EdoRguez/business-manager/gateway/pkg/customer/contracts"
	"github.com/EdoRguez/business-manager/gateway/pkg/customer/pb"
	"github.com/EdoRguez/business-manager/gateway/pkg/util/query_params"
)

func GetCustomers(w http.ResponseWriter, r *http.Request, c *config.Config) {
	w.Header().Set("Content-Type", "application/json")
	companyId := query_params.GetId("companyId", r)
	firstName := r.URL.Query().Get("name")
	lastName := r.URL.Query().Get("lastName")
	identificationNumber := r.URL.Query().Get("identificationNumber")
	limit, offset := query_params.GetFilter(r)

	fmt.Println("hola")
	fmt.Println(firstName)

	params := &pb.GetCustomersRequest{
		CompanyId:            companyId,
		FirstName:            &firstName,
		LastName:             &lastName,
		IdentificationNumber: &identificationNumber,
		Limit:                limit,
		Offset:               offset,
	}

	if err := client.InitCustomerServiceClient(c); err != nil {
		json.NewEncoder(w).Encode(&contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	res, err := client.GetCustomers(params, r.Context())

	if err != nil {
		fmt.Println("API Gateway :  GetCustomers - ERROR")
		json.NewEncoder(w).Encode(err)
		return

	}

	fmt.Println("API Gateway :  GetCustomers - SUCCESS")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
