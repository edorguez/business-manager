package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/EdoRguez/business-manager/gateway/pkg/customer/pb"
	"github.com/gorilla/mux"
)

func GetCustomer(w http.ResponseWriter, r *http.Request, c pb.CustomerServiceClient) {
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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}
