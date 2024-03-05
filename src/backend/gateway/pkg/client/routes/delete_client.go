package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/EdoRguez/business-manager/gateway/pkg/client/pb"
	"github.com/gorilla/mux"
)

func DeleteClient(w http.ResponseWriter, r *http.Request, c pb.ClientServiceClient) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert Id", http.StatusBadRequest)
	}

	params := &pb.DeleteClientRequest{
		Id: int64(id),
	}

	res, err := c.DeleteClient(r.Context(), params)

	if err != nil {
		fmt.Println("API Gateway :  DeleteClient - ERROR")
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("API Gateway :  DeleteClient - SUCCESS")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}
