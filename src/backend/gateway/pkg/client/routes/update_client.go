package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/EdoRguez/business-manager/gateway/pkg/client/pb"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type UpdateClientRequestBody struct {
	FirstName            string `json:"firstName" validate:"required,max=20"`
	LastName             string `json:"lastName" validate:"max=20"`
	Email                string `json:"email" validate:"email,max=100"`
	Phone                string `json:"phone" validate:"max=11"`
	IdentificationNumber string `json:"identificationNumber" validate:"required,max=20"`
	IdentificationType   string `json:"identificationType" validate:"required,max=1"`
}

func (c *UpdateClientRequestBody) Validate() error {
	validate := validator.New()

	return validate.Struct(c)
}

func UpdateClient(w http.ResponseWriter, r *http.Request, c pb.ClientServiceClient) {
	fmt.Println("API Gateway :  UpdateClient")

	// We got our body through context, since we saved it in a middleware
	body := r.Context().Value(UpdateClientRequestBody{}).(UpdateClientRequestBody)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert Id", http.StatusBadRequest)
	}

	fmt.Println("API Gateway :  UpdateClient - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	updateClientParams := &pb.UpdateClientRequest{
		Id:                   int64(id),
		FirstName:            body.FirstName,
		LastName:             body.LastName,
		Email:                body.Email,
		Phone:                body.Phone,
		IdentificationNumber: body.IdentificationNumber,
		IdentificationType:   body.IdentificationType,
	}

	res, err := c.UpdateClient(r.Context(), updateClientParams)

	if err != nil {
		fmt.Println("API Gateway :  UpdateClient - ERROR")
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("API Gateway :  UpdateClient - SUCCESS")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}
