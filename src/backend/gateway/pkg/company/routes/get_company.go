package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/EdoRguez/business-manager/gateway/pkg/company/contracts"
	"github.com/EdoRguez/business-manager/gateway/pkg/company/pb"
	"github.com/gorilla/mux"
)

func GetCompany(w http.ResponseWriter, r *http.Request, c pb.CompanyServiceClient) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert Id", http.StatusBadRequest)
	}

	params := &pb.GetCompanyRequest{
		Id: int64(id),
	}

	res, err := c.GetCompany(r.Context(), params)

	if err != nil {
		fmt.Println("API Gateway :  GetCompany - ERROR")
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("API Gateway :  GetCompany - SUCCESS")
	w.WriteHeader(int(res.Status))

	if res.Status != http.StatusOK {
		json.NewEncoder(w).Encode(contracts.Error{
			Status: res.Status,
			Error:  res.Error,
		})
		return
	}

	json.NewEncoder(w).Encode(&contracts.GetCompanyResponse{
		Id:       res.Id,
		Name:     res.Name,
		ImageUrl: res.ImageUrl,
	})
}
