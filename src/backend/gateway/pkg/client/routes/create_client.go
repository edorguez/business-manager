package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/gateway/pkg/client/pb"
	"github.com/go-playground/validator/v10"
)

type CreateClientRequestBody struct {
	CompanyId            int64  `json:"companyId" validate:"required"`
	FirstName            string `json:"firstName" validate:"required,max=20"`
	LastName             string `json:"lastName" validate:"max=20"`
	Email                string `json:"email" validate:"email,max=100"`
	Phone                string `json:"phone" validate:"max=11"`
	IdentificationNumber string `json:"identificationNumber" validate:"required,max=20"`
	IdentificationType   string `json:"identificationType" validate:"required,max=1"`
}

func (c *CreateClientRequestBody) Validate() error {
	validate := validator.New()

	return validate.Struct(c)
}

func CreateClient(w http.ResponseWriter, r *http.Request, c pb.ClientServiceClient) {
	fmt.Println("API Gateway :  CreateClient")

	// We got out body through context since we saved it in a middleware
	body := r.Context().Value(CreateClientRequestBody{}).(CreateClientRequestBody)

	fmt.Println("API Gateway :  CreateClient - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	createClientParams := &pb.CreateClientRequest{
		CompanyId:            body.CompanyId,
		FirstName:            body.FirstName,
		LastName:             body.LastName,
		Email:                body.Email,
		Phone:                body.Phone,
		IdentificationNumber: body.IdentificationNumber,
		IdentificationType:   body.IdentificationType,
	}

	res, err := c.CreateClient(r.Context(), createClientParams)

	if err != nil {
		fmt.Println("API Gateway :  CreateClient - ERROR")
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("API Gateway :  CreateClient - SUCCESS")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}
