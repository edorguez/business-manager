package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/EdoRguez/business-manager/gateway/pkg/company/pb"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type UpdateCompanyRequestBody struct {
	Name     string `json:"name" validate:"required,max=50"`
	ImageUrl string `json:"imageUrl"  validate:"required"`
}

func (c *UpdateCompanyRequestBody) Validate() error {
	validate := validator.New()

	return validate.Struct(c)
}

func UpdateCompany(w http.ResponseWriter, r *http.Request, c pb.CompanyServiceClient) {
	fmt.Println("API Gateway :  UpdateCompany")

	// We got our body through context, since we saved it in a middleware
	body := r.Context().Value(UpdateCompanyRequestBody{}).(UpdateCompanyRequestBody)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert Id", http.StatusBadRequest)
	}

	fmt.Println("API Gateway :  UpdateCompany - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	updateCompanyParams := &pb.UpdateCompanyRequest{
		Id:       int64(id),
		Name:     body.Name,
		ImageUrl: body.ImageUrl,
	}

	res, err := c.UpdateCompany(r.Context(), updateCompanyParams)

	if err != nil {
		fmt.Println("API Gateway :  UpdateCompany - ERROR")
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("API Gateway :  UpdateCompany - SUCCESS")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}
