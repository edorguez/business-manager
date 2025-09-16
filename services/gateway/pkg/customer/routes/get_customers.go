package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/edorguez/business-manager/services/gateway/pkg/config"
	"github.com/edorguez/business-manager/services/gateway/pkg/customer/client"
	pb "github.com/edorguez/business-manager/shared/pb/customer"
	"github.com/edorguez/business-manager/shared/types"
	"github.com/edorguez/business-manager/shared/util/query_params"
)

func GetCustomers(w http.ResponseWriter, r *http.Request, c *config.Config) {
	w.Header().Set("Content-Type", "application/json")
	companyId := query_params.GetId("companyId", r)
	firstName := r.URL.Query().Get("name")
	lastName := r.URL.Query().Get("lastName")
	identificationNumber := r.URL.Query().Get("identificationNumber")
	limit, offset := query_params.GetFilter(r)

	params := &pb.GetCustomersRequest{
		CompanyId:            companyId,
		FirstName:            &firstName,
		LastName:             &lastName,
		IdentificationNumber: &identificationNumber,
		Limit:                limit,
		Offset:               offset,
	}

	if err := client.InitCustomerServiceClient(c); err != nil {
		json.NewEncoder(w).Encode(&types.Error{
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
