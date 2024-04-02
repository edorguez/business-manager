package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/gateway/pkg/company/pb"
	"github.com/EdoRguez/business-manager/gateway/pkg/util"
)

func GetPayments(w http.ResponseWriter, r *http.Request, c pb.PaymentServiceClient) {
	companyId := util.GetIdQueryParam("companyId", r)
	limit, offset := util.GetFilterQueryParams(r)

	params := &pb.GetPaymentsRequest{
		CompanyId: companyId,
		Limit:     limit,
		Offset:    offset,
	}

	res, err := c.GetPayments(r.Context(), params)

	if err != nil {
		fmt.Println("API Gateway :  GetPayments - ERROR")
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("API Gateway :  GetPayments - SUCCESS")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}
