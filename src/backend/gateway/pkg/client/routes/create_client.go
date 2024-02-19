package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/gateway/pkg/client/pb"
)

type CreateClientRequestBody struct {
	CompanyId            int64  `json:"companyId"`
	FirstName            string `json:"firstName"`
	LastName             string `json:"lastName"`
	Email                string `json:"email"`
	Phone                string `json:"phone"`
	IdentificationNumber string `json:"identificationNumber"`
	IdentificationType   string `json:"identificationType"`
}

func CreateClient(w http.ResponseWriter, r *http.Request, c pb.ClientServiceClient) {
	fmt.Println("API Gateway :  Create Client")
	var body CreateClientRequestBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println("api gateway")
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

	_, err := c.CreateClient(r.Context(), createClientParams)

	if err != nil {
		http.Error(w, "There is an error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
