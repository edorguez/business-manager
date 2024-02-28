package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/gateway/pkg/client/pb"
)

func GetClients(w http.ResponseWriter, r *http.Request, c pb.ClientServiceClient) {

	params := &pb.GetClientsRequest{
		Limit:  10,
		Offset: 0,
	}

	res, err := c.GetClients(r.Context(), params)

	if err != nil {
		fmt.Println("API Gateway :  GetClient - ERROR")
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("API Gateway :  GetClient - SUCCESS")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}
