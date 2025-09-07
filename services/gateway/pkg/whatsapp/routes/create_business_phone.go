package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/edorguez/business-manager/services/gateway/pkg/config"
	"github.com/edorguez/business-manager/services/gateway/pkg/whatsapp/client"
	"github.com/edorguez/business-manager/services/gateway/pkg/whatsapp/contracts"
)

func CreateBusinessPhone(w http.ResponseWriter, r *http.Request, c *config.Config) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("API Gateway :  CreateBusinessPhone")

	// We got our body through context, since we saved it in a middleware
	body := r.Context().Value(contracts.CreateBusinessPhoneRequest{}).(contracts.CreateBusinessPhoneRequest)

	fmt.Println("API Gateway :  CreateBusinessPhone - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	if err := client.InitWhatsappServiceClient(c); err != nil {
		json.NewEncoder(w).Encode(&contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	res, errCreate := client.CreateBusinessPhone(body, r.Context())

	if errCreate != nil {
		fmt.Println("API Gateway :  CreateBusinessPhone - ERROR")
		json.NewEncoder(w).Encode(errCreate)
		return
	}

	fmt.Println("API Gateway :  CreateBusinessPhone - SUCCESS")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}
