package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/EdoRguez/business-manager/gateway/pkg/customer/contracts"
	"github.com/EdoRguez/business-manager/gateway/pkg/customer/pb"
	"github.com/gorilla/mux"
)

func UpdateCustomer(w http.ResponseWriter, r *http.Request, c pb.CustomerServiceClient) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("API Gateway :  UpdateCustomer")

	// We got our body through context, since we saved it in a middleware
	body := r.Context().Value(contracts.UpdateCustomerRequest{}).(contracts.UpdateCustomerRequest)

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
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}
