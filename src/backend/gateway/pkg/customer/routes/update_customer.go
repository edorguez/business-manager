package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/EdoRguez/business-manager/gateway/pkg/customer/pb"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type UpdateCustomerRequestBody struct {
	FirstName            string `json:"firstName" validate:"required,max=20"`
	LastName             string `json:"lastName" validate:"max=20"`
	Email                string `json:"email" validate:"email,max=100"`
	Phone                string `json:"phone" validate:"max=11"`
	IdentificationNumber string `json:"identificationNumber" validate:"required,max=20"`
	IdentificationType   string `json:"identificationType" validate:"required,max=1"`
}

func (c *UpdateCustomerRequestBody) Validate() error {
	validate := validator.New()

	return validate.Struct(c)
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request, c pb.CustomerServiceClient) {
	fmt.Println("API Gateway :  UpdateCustomer")

	// We got our body through context, since we saved it in a middleware
	body := r.Context().Value(UpdateCustomerRequestBody{}).(UpdateCustomerRequestBody)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert Id", http.StatusBadRequest)
	}

	fmt.Println("API Gateway :  UpdateCustomer - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	updateCustomerParams := &pb.UpdateCustomerRequest{
		Id:                   int64(id),
		FirstName:            body.FirstName,
		LastName:             body.LastName,
		Email:                body.Email,
		Phone:                body.Phone,
		IdentificationNumber: body.IdentificationNumber,
		IdentificationType:   body.IdentificationType,
	}

	res, err := c.UpdateCustomer(r.Context(), updateCustomerParams)

	if err != nil {
		fmt.Println("API Gateway :  UpdateCustomer - ERROR")
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("API Gateway :  UpdateCustomer - SUCCESS")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}
