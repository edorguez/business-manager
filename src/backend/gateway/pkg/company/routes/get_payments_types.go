package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/gateway/pkg/company/client"
	"github.com/EdoRguez/business-manager/gateway/pkg/company/contracts"
	"github.com/EdoRguez/business-manager/gateway/pkg/company/pb"
	"github.com/EdoRguez/business-manager/gateway/pkg/config"
	"github.com/EdoRguez/business-manager/gateway/pkg/util/query_params"
)

func GetPaymentsTypes(w http.ResponseWriter, r *http.Request, c *config.Config) {
	w.Header().Set("Content-Type", "application/json")
	companyId := query_params.GetId("companyId", r)

	if companyId <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(contracts.Error{
			Status: http.StatusBadRequest,
			Error:  "Company ID is required in order to get results",
		})
		return
	}

	params := &pb.GetPaymentsTypesRequest{
		CompanyId: companyId,
	}

	if err := client.InitPaymentServiceClient(c); err != nil {
		json.NewEncoder(w).Encode(&contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	res, err := client.GetPaymentsTypes(params, r.Context())

	if err != nil {
		fmt.Println("API Gateway :  GetPaymentsTypes - ERROR")
		json.NewEncoder(w).Encode(err)
		return
	}

	fmt.Println("API Gateway :  GetPaymentsTypes - SUCCESS")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
