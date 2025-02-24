package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/gateway/pkg/config"
	"github.com/EdoRguez/business-manager/gateway/pkg/whatsapp/client"
	"github.com/EdoRguez/business-manager/gateway/pkg/whatsapp/contracts"
)

func UpdateBusinessNumber(w http.ResponseWriter, r *http.Request, c *config.Config) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("API Gateway :  UpdateBusinessNumber")

	// We got our body through context, since we saved it in a middleware
	body := r.Context().Value(contracts.UpdateBusinessPhoneRequest{}).(contracts.UpdateBusinessPhoneRequest)

	fmt.Println("API Gateway :  UpdateBusinessNumber - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	if err := client.InitWhatsappServiceClient(c); err != nil {
		json.NewEncoder(w).Encode(&contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	res, errUpdate := client.UpdateBusinessPhone(body, r.Context())

	if errUpdate != nil {
		fmt.Println("API Gateway :  UpdateBusinessNumber - ERROR")
		json.NewEncoder(w).Encode(errUpdate)
		return
	}

	fmt.Println("API Gateway :  UpdateBusinessNumber - SUCCESS")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}
