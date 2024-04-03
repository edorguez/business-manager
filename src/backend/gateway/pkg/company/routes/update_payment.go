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

func UpdatePayment(w http.ResponseWriter, r *http.Request, c pb.PaymentServiceClient) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("API Gateway :  UpdatePayment")

	// We got our body through context, since we saved it in a middleware
	body := r.Context().Value(contracts.UpdatePaymentRequest{}).(contracts.UpdatePaymentRequest)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert Id", http.StatusBadRequest)
	}

	fmt.Println("API Gateway :  UpdatePayment - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	updatePaymentParams := &pb.UpdatePaymentRequest{
		Id:                   int64(id),
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

	res, err := c.UpdatePayment(r.Context(), updatePaymentParams)

	if err != nil {
		fmt.Println("API Gateway :  UpdatePayment - ERROR")
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("API Gateway :  UpdatePayment - SUCCESS")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}
