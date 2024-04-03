package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/gateway/pkg/company/contracts"
	"github.com/EdoRguez/business-manager/gateway/pkg/company/pb"
)

func CreateCompany(w http.ResponseWriter, r *http.Request, c pb.CompanyServiceClient) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("API Gateway :  CreateCompany")

	// We got our body through context, since we saved it in a middleware
	body := r.Context().Value(contracts.CreateCompanyRequest{}).(contracts.CreateCompanyRequest)

	fmt.Println("API Gateway :  CreateCompany - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	createCompanyParams := &pb.CreateCompanyRequest{
		Name:     body.Name,
		ImageUrl: body.ImageUrl,
	}

	res, err := c.CreateCompany(r.Context(), createCompanyParams)

	if err != nil {
		fmt.Println("API Gateway :  CreateCompany - ERROR")
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("API Gateway :  CreateCompany - SUCCESS")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}
