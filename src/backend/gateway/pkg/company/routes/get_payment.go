package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/EdoRguez/business-manager/gateway/pkg/company/contracts"
	"github.com/EdoRguez/business-manager/gateway/pkg/company/pb"
	"github.com/gorilla/mux"
)

func GetPayment(w http.ResponseWriter, r *http.Request, c pb.PaymentServiceClient) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert Id", http.StatusBadRequest)
	}

	params := &pb.GetPaymentRequest{
		Id: int64(id),
	}

	res, err := c.GetPayment(r.Context(), params)

	if err != nil {
		fmt.Println("API Gateway :  GetPayment - ERROR")
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("API Gateway :  GetPayment - SUCCESS")
	w.WriteHeader(int(res.Status))

	if res.Status != http.StatusOK {
		json.NewEncoder(w).Encode(contracts.Error{
			Status: res.Status,
			Error:  res.Error,
		})
		return
	}

	json.NewEncoder(w).Encode(contracts.GetPaymentResponse{
		Id:                   res.Id,
		CompanyId:            res.CompanyId,
		Name:                 res.Name,
		Bank:                 res.Bank,
		AccountNumber:        res.AccountNumber,
		AccountType:          res.AccountType,
		IdentificationNumber: res.IdentificationNumber,
		IdentificationType:   res.IdentificationType,
		Phone:                res.Phone,
		Email:                res.Email,
		PaymentTypeId:        res.PaymentTypeId,
		PaymentType: &contracts.GetPaymentTypeResponse{
			Id:   res.PaymentType.Id,
			Name: res.PaymentType.Name,
		},
	})
}
