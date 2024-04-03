package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/EdoRguez/business-manager/gateway/pkg/company/pb"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type UpdatePaymentRequestBody struct {
	Name                 string  `json:"name" validate:"required,max=50"`
	Bank                 *string `json:"bank" validate:"omitempty,max=50"`
	AccountNumber        *string `json:"accountNumber" validate:"omitempty,max=20"`
	AccountType          *string `json:"accountType" validate:"omitempty,max=20"`
	IdentificationNumber *string `json:"identificationNumber" validate:"omitempty,max=20"`
	IdentificationType   *string `json:"identificationType" validate:"omitempty,max=1"`
	Phone                *string `json:"phone" validate:"omitempty,max=11"`
	Email                *string `json:"email" validate:"omitempty,max=100"`
	PaymentTypeId        int64   `json:"paymentTypeId" validate:"required"`
}

func (c *UpdatePaymentRequestBody) Validate() error {
	validate := validator.New()

	return validate.Struct(c)
}

func UpdatePayment(w http.ResponseWriter, r *http.Request, c pb.PaymentServiceClient) {
	fmt.Println("API Gateway :  UpdatePayment")

	// We got our body through context, since we saved it in a middleware
	body := r.Context().Value(UpdatePaymentRequestBody{}).(UpdatePaymentRequestBody)

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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}
