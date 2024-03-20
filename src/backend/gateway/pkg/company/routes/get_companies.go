package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/gateway/pkg/company/pb"
	"github.com/EdoRguez/business-manager/gateway/pkg/util"
)

func GetCompanies(w http.ResponseWriter, r *http.Request, c pb.CompanyServiceClient) {
	limit, offset := util.GetFilterQueryParams(r)

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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}
