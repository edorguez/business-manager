package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/EdoRguez/business-manager/gateway/pkg/client/pb"
	"github.com/gorilla/mux"
)

type UpdateClientRequestBody struct {
	FirstName            string `json:"firstName" binding:"required,max=20"`
	LastName             string `json:"lastName" binding:"max=20"`
	Email                string `json:"email" binding:"email,max=100"`
	Phone                string `json:"phone" binding:"max=11"`
	IdentificationNumber string `json:"identificationNumber" binding:"required,max=20"`
	IdentificationType   string `json:"identificationType" binding:"required,max=1"`
}

func UpdateClient(w http.ResponseWriter, r *http.Request, c pb.ClientServiceClient) {
	fmt.Println("API Gateway :  UpdateClient")
	var body UpdateClientRequestBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

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
