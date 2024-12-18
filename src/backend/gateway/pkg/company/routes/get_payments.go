package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/gateway/pkg/company/client"
	"github.com/EdoRguez/business-manager/gateway/pkg/company/contracts"
	"github.com/EdoRguez/business-manager/gateway/pkg/config"
	pb "github.com/EdoRguez/business-manager/gateway/pkg/pb/payment"
	"github.com/EdoRguez/business-manager/gateway/pkg/util/query_params"
)

func GetPayments(w http.ResponseWriter, r *http.Request, c *config.Config) {
	w.Header().Set("Content-Type", "application/json")
	companyId := query_params.GetId("companyId", r)
	paymentTypeId := query_params.GetId("paymentTypeId", r)
	limit, offset := query_params.GetFilter(r)

	if companyId <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(contracts.Error{
			Status: http.StatusBadRequest,
			Error:  "Company ID is required in order to get results",
		})
		return
	}

	if paymentTypeId < 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(contracts.Error{
			Status: http.StatusBadRequest,
			Error:  "Payment Type ID is required in order to get results",
		})
		return
	}

	params := &pb.GetPaymentsRequest{
		CompanyId:     companyId,
		PaymentTypeId: paymentTypeId,
		Limit:         limit,
		Offset:        offset,
	}

	if err := client.InitPaymentServiceClient(c); err != nil {
		json.NewEncoder(w).Encode(&contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	res, err := client.GetPayments(params, r.Context())

	if err != nil {
		fmt.Println("API Gateway :  GetPayments - ERROR")
		json.NewEncoder(w).Encode(err)
		return
	}

	fmt.Println("API Gateway :  GetPayments - SUCCESS")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
