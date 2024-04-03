package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/gateway/pkg/customer/contracts"
	"github.com/EdoRguez/business-manager/gateway/pkg/customer/pb"
)

func CreateCustomer(w http.ResponseWriter, r *http.Request, c pb.CustomerServiceClient) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("API Gateway :  CreateCustomer")

	// We got our body through context, since we saved it in a middleware
	body := r.Context().Value(contracts.CreateCustomerRequest{}).(contracts.CreateCustomerRequest)

	fmt.Println("API Gateway :  CreateCustomer - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	createCustomerParams := &pb.CreateCustomerRequest{
		CompanyId:            body.CompanyId,
		FirstName:            body.FirstName,
		LastName:             body.LastName,
		Email:                body.Email,
		Phone:                body.Phone,
		IdentificationNumber: body.IdentificationNumber,
		IdentificationType:   body.IdentificationType,
	}

	res, err := c.CreateCustomer(r.Context(), createCustomerParams)

	if err != nil {
		fmt.Println("API Gateway :  CreateCustomer - ERROR")
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("API Gateway :  CreateCustomer - SUCCESS")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}
