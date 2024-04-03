package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/gateway/pkg/company/pb"
	"github.com/go-playground/validator/v10"
)

type CreateCompanyRequestBody struct {
	Name     string  `json:"name" validate:"required,max=50"`
	ImageUrl *string `json:"imageUrl" validate:"omitempty,required"`
}

func (c *CreateCompanyRequestBody) Validate() error {
	validate := validator.New()

	return validate.Struct(c)
}

func CreateCompany(w http.ResponseWriter, r *http.Request, c pb.CompanyServiceClient) {
	fmt.Println("API Gateway :  CreateCompany")

	// We got our body through context, since we saved it in a middleware
	body := r.Context().Value(CreateCompanyRequestBody{}).(CreateCompanyRequestBody)

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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}
