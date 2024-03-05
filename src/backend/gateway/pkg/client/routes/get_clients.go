package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/gateway/pkg/client/pb"
	"github.com/EdoRguez/business-manager/gateway/pkg/util"
)

func GetClients(w http.ResponseWriter, r *http.Request, c pb.ClientServiceClient) {
	companyId := util.GetIdQueryParam("companyId", r)
	limit, offset := util.GetFilterQueryParams(r)

	params := &pb.GetClientsRequest{
		CompanyId: companyId,
		Limit:     limit,
		Offset:    offset,
	}

	res, err := c.GetClients(r.Context(), params)

	if err != nil {
		fmt.Println("API Gateway :  GetClient - ERROR")
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("API Gateway :  GetClient - SUCCESS")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}
