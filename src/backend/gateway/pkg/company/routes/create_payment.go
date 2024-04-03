package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/gateway/pkg/company/contracts"
	"github.com/EdoRguez/business-manager/gateway/pkg/company/pb"
)

func CreatePayment(w http.ResponseWriter, r *http.Request, c pb.PaymentServiceClient) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("API Gateway :  CreatePayment")

	// We got our body through context, since we saved it in a middleware
	body := r.Context().Value(contracts.CreatePaymentRequest{}).(contracts.CreatePaymentRequest)

	fmt.Println("API Gateway :  CreatePayment - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	createPaymentParams := &pb.CreatePaymentRequest{
		CompanyId:            body.CompanyId,
		Name:                 body.Name,
		Bank:                 body.Bank,
		AccountNumber:        body.AccountNumber,
		AccountType:          body.AccountType,
		IdentificationNumber: body.IdentificationNumber,
		IdentificationType:   body.IdentificationType,
		Phone:                body.Phone,
		Email:                body.Email,
		PaymentTypeId:        body.PaymentTypeId,
	}

	res, err := c.CreatePayment(r.Context(), createPaymentParams)

	if err != nil {
		fmt.Println("API Gateway :  CreatePayment - ERROR")
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("API Gateway :  CreatePayment - SUCCESS")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}
