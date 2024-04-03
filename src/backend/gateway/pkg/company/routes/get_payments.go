package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/gateway/pkg/company/contracts"
	"github.com/EdoRguez/business-manager/gateway/pkg/company/pb"
	"github.com/EdoRguez/business-manager/gateway/pkg/util/query_params"
)

func GetPayments(w http.ResponseWriter, r *http.Request, c pb.PaymentServiceClient) {
	w.Header().Set("Content-Type", "application/json")
	companyId := query_params.GetId("companyId", r)
	limit, offset := query_params.GetFilter(r)

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
	w.WriteHeader(int(res.Status))

	if res.Status != http.StatusOK {
		json.NewEncoder(w).Encode(contracts.Error{
			Status: res.Status,
			Error:  res.Error,
		})
		return
	}

	var pr []*contracts.GetPaymentResponse
	for _, v := range res.Payments {
		pr = append(pr, &contracts.GetPaymentResponse{
			Id:                   v.Id,
			CompanyId:            v.CompanyId,
			Name:                 v.Name,
			Bank:                 v.Bank,
			AccountNumber:        v.Bank,
			AccountType:          v.AccountType,
			IdentificationNumber: v.IdentificationNumber,
			IdentificationType:   v.IdentificationType,
			Phone:                v.Phone,
			Email:                v.Email,
			PaymentTypeId:        v.PaymentTypeId,
			PaymentType: &contracts.GetPaymentTypeResponse{
				Id:   v.PaymentType.Id,
				Name: v.PaymentType.Name,
			},
		})
	}

	json.NewEncoder(w).Encode(pr)
}
