package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/gateway/pkg/company/pb"
	"github.com/go-playground/validator/v10"
)

type CreatePaymentRequestBody struct {
	CompanyId            int64  `json:"companyId" validate:"required"`
	Name                 string `json:"name" validate:"required,max=50"`
	Bank                 string `json:"bank" validate:"max=50"`
	AccountNumber        string `json:"accountNumber" validate:"max=20"`
	AccountType          string `json:"accountType" validate:"max=20"`
	IdentificationNumber string `json:"identificationNumber" validate:"max=20"`
	IdentificationType   string `json:"identificationType" validate:"max=1"`
	Phone                string `json:"phone" validate:"max=11"`
	Email                string `json:"email" validate:"max=100"`
	PaymentTypeId        int64  `json:"paymentTypeId" validate:"required"`
}

func (c *CreatePaymentRequestBody) Validate() error {
	validate := validator.New()

	return validate.Struct(c)
}

func CreatePayment(w http.ResponseWriter, r *http.Request, c pb.PaymentServiceClient) {
	fmt.Println("API Gateway :  CreatePayment")

	// We got our body through context, since we saved it in a middleware
	body := r.Context().Value(CreatePaymentRequestBody{}).(CreatePaymentRequestBody)

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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}
