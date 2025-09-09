package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/edorguez/business-manager/services/gateway/pkg/company/client"
	"github.com/edorguez/business-manager/services/gateway/pkg/config"
	pb "github.com/edorguez/business-manager/shared/pb/payment_type"
	"github.com/edorguez/business-manager/shared/types"
	"github.com/edorguez/business-manager/shared/util/query_params"
)

func GetPaymentTypes(w http.ResponseWriter, r *http.Request, c *config.Config) {
	w.Header().Set("Content-Type", "application/json")
	limit, offset := query_params.GetFilter(r)

	params := &pb.GetPaymentTypesRequest{
		Limit:  limit,
		Offset: offset,
	}

	if err := client.InitPaymentTypeServiceClient(c); err != nil {
		json.NewEncoder(w).Encode(&types.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	res, err := client.GetPaymentTypes(params, r.Context())

	if err != nil {
		fmt.Println("API Gateway :  GetPaymentTypes - ERROR")
		json.NewEncoder(w).Encode(err)
		return
	}

	fmt.Println("API Gateway :  GetPaymentTypes - SUCCESS")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
