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

func GetCustomer(w http.ResponseWriter, r *http.Request, c pb.CustomerServiceClient) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert Id", http.StatusBadRequest)
	}

	params := &pb.GetCustomerRequest{
		Id: int64(id),
	}

	res, err := c.GetCustomer(r.Context(), params)

	if err != nil {
		fmt.Println("API Gateway :  GetCustomer - ERROR")
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("API Gateway :  GetCustomer - SUCCESS")
	w.WriteHeader(int(res.Status))

	if res.Status != http.StatusOK {
		json.NewEncoder(w).Encode(contracts.Error{
			Status: res.Status,
			Error:  res.Error,
		})
		return
	}

	json.NewEncoder(w).Encode(&contracts.GetCustomerResponse{
		Id:                   res.Id,
		CompanyId:            res.CompanyId,
		FirstName:            res.FirstName,
		LastName:             res.LastName,
		Email:                res.Email,
		Phone:                res.Phone,
		IdentificationNumber: res.IdentificationNumber,
		IdentificationType:   res.IdentificationType,
	})
}
