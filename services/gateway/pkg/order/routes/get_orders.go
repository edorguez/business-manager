package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	companyClient "github.com/edorguez/business-manager/services/gateway/pkg/company/client"
	"github.com/edorguez/business-manager/services/gateway/pkg/config"
	orderClient "github.com/edorguez/business-manager/services/gateway/pkg/order/client"
	"github.com/edorguez/business-manager/shared/types"
)

func GetOrders(w http.ResponseWriter, r *http.Request, c *config.Config) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("API Gateway :  GetOrders")

	// Parse query parameters
	companyIdStr := r.URL.Query().Get("companyId")
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	if companyIdStr == "" {
		json.NewEncoder(w).Encode(&types.Error{
			Status: http.StatusBadRequest,
			Error:  "Company ID is required",
		})
		return
	}

	companyId, err := strconv.ParseInt(companyIdStr, 10, 64)
	if err != nil {
		json.NewEncoder(w).Encode(&types.Error{
			Status: http.StatusBadRequest,
			Error:  "Invalid company ID",
		})
		return
	}

	limit := int32(10) // default
	if limitStr != "" {
		l, err := strconv.ParseInt(limitStr, 10, 32)
		if err == nil && l > 0 && l <= 100 {
			limit = int32(l)
		}
	}

	offset := int32(0) // default
	if offsetStr != "" {
		o, err := strconv.ParseInt(offsetStr, 10, 32)
		if err == nil && o >= 0 {
			offset = int32(o)
		}
	}

	if err := orderClient.InitOrderServiceClient(c); err != nil {
		json.NewEncoder(w).Encode(&types.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	if err := companyClient.InitCompanyServiceClient(c); err != nil {
		json.NewEncoder(w).Encode(&types.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	// Validate company exists
	_, errCompany := companyClient.GetCompany(companyId, r.Context())
	if errCompany != nil {
		fmt.Println("API Gateway :  GetOrders - ERROR")
		errCompany.Error = "Company not found"
		json.NewEncoder(w).Encode(errCompany)
		return
	}

	res, errResp := orderClient.GetOrders(companyId, limit, offset, r.Context())
	if errResp != nil {
		fmt.Println("API Gateway :  GetOrders - ERROR")
		json.NewEncoder(w).Encode(errResp)
		return
	}

	fmt.Println("API Gateway :  GetOrders - SUCCESS")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}
