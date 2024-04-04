package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/gateway/pkg/customer/contracts"
	"github.com/EdoRguez/business-manager/gateway/pkg/customer/pb"
	"github.com/EdoRguez/business-manager/gateway/pkg/util/query_params"
)

func GetCustomers(w http.ResponseWriter, r *http.Request, c pb.CustomerServiceClient) {
	w.Header().Set("Content-Type", "application/json")
	companyId := query_params.GetId("companyId", r)
	limit, offset := query_params.GetFilter(r)

	if companyId <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(contracts.Error{
			Status: http.StatusBadRequest,
			Error:  "Company ID is required in order to get results",
		})
		return
	}

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
	w.WriteHeader(int(res.Status))

	if res.Status != http.StatusOK {
		json.NewEncoder(w).Encode(contracts.Error{
			Status: res.Status,
			Error:  res.Error,
		})
		return
	}

	cr := make([]*contracts.GetCustomerResponse, 0)
	for _, v := range res.Customers {
		cr = append(cr, &contracts.GetCustomerResponse{
			Id:                   v.Id,
			CompanyId:            v.CompanyId,
			FirstName:            v.FirstName,
			LastName:             v.LastName,
			Email:                v.Email,
			Phone:                v.Phone,
			IdentificationNumber: v.IdentificationNumber,
			IdentificationType:   v.IdentificationType,
		})
	}

	json.NewEncoder(w).Encode(cr)
}
