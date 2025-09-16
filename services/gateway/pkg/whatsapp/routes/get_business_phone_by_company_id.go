package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/edorguez/business-manager/services/gateway/pkg/config"
	"github.com/edorguez/business-manager/services/gateway/pkg/whatsapp/client"
	"github.com/edorguez/business-manager/shared/types"
	"github.com/gorilla/mux"
)

func GetBusinessPhoneByCompanyId(w http.ResponseWriter, r *http.Request, c *config.Config) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		json.NewEncoder(w).Encode(&types.Error{
			Status: http.StatusBadRequest,
			Error:  "Unable to convert ID",
		})
		return
	}

	if err := client.InitWhatsappServiceClient(c); err != nil {
		json.NewEncoder(w).Encode(&types.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	res, errPhone := client.GetBusinessPhoneByCompanyId(int64(id), r.Context())

	if errPhone != nil {
		fmt.Println("API Gateway :  GetBusinessPhoneByCompanyId - ERROR")
		json.NewEncoder(w).Encode(errPhone)
		return
	}

	fmt.Println("API Gateway :  GetBusinessPhoneByCompanyId - SUCCESS")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
