package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/gateway/pkg/client/pb"
)

type CreateClientRequestBody struct {
	CompanyId            int64  `json:"companyId" binding:"required"`
	FirstName            string `json:"firstName" binding:"required,max=20"`
	LastName             string `json:"lastName" binding:"max=20"`
	Email                string `json:"email" binding:"email,max=100"`
	Phone                string `json:"phone" binding:"max=11"`
	IdentificationNumber string `json:"identificationNumber" binding:"required,max=20"`
	IdentificationType   string `json:"identificationType" binding:"required,max=1"`
}

func CreateClient(w http.ResponseWriter, r *http.Request, c pb.ClientServiceClient) {
	fmt.Println("API Gateway :  CreateClient")
	var body CreateClientRequestBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}
