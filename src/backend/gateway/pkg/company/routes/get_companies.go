package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/gateway/pkg/company/contracts"
	"github.com/EdoRguez/business-manager/gateway/pkg/company/pb"
	"github.com/EdoRguez/business-manager/gateway/pkg/util/query_params"
)

func GetCompanies(w http.ResponseWriter, r *http.Request, c pb.CompanyServiceClient) {
	w.Header().Set("Content-Type", "application/json")
	limit, offset := query_params.GetFilter(r)

	params := &pb.GetCompaniesRequest{
		Limit:  limit,
		Offset: offset,
	}

	res, err := c.GetCompanies(r.Context(), params)

	if err != nil {
		fmt.Println("API Gateway :  GetCompanies - ERROR")
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("API Gateway :  GetCompanies - SUCCESS")
	w.WriteHeader(int(res.Status))

	if res.Status != http.StatusOK {
		json.NewEncoder(w).Encode(contracts.Error{
			Status: res.Status,
			Error:  res.Error,
		})
		return
	}

	var cr []*contracts.GetCompanyResponse
	for _, v := range res.Companies {
		cr = append(cr, &contracts.GetCompanyResponse{
			Id:       v.Id,
			Name:     v.Name,
			ImageUrl: v.ImageUrl,
		})
	}

	json.NewEncoder(w).Encode(cr)
}
