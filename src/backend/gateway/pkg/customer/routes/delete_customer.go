package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/EdoRguez/business-manager/gateway/pkg/customer/pb"
	"github.com/gorilla/mux"
)

func DeleteCustomer(w http.ResponseWriter, r *http.Request, c pb.CustomerServiceClient) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert Id", http.StatusBadRequest)
	}

	params := &pb.DeleteCustomerRequest{
		Id: int64(id),
	}

	res, err := c.DeleteCustomer(r.Context(), params)

	if err != nil {
		fmt.Println("API Gateway :  DeleteCustomer - ERROR")
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("API Gateway :  DeleteCustomer - SUCCESS")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}
