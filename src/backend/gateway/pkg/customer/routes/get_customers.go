package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/gateway/pkg/customer/pb"
	"github.com/EdoRguez/business-manager/gateway/pkg/util"
)

func GetCustomers(w http.ResponseWriter, r *http.Request, c pb.CustomerServiceClient) {
	companyId := util.GetIdQueryParam("companyId", r)
	limit, offset := util.GetFilterQueryParams(r)

	params := &pb.GetCustomersRequest{
		CompanyId: companyId,
		Limit:     limit,
		Offset:    offset,
	}

	res, err := c.GetCustomers(r.Context(), params)

	if err != nil {
		fmt.Println("API Gateway :  GetCustomers - ERROR")
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("API Gateway :  GetCustomers - SUCCESS")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}
