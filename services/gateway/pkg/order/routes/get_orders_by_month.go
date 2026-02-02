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

func GetOrdersByMonth(w http.ResponseWriter, r *http.Request, c *config.Config) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("API Gateway :  GetOrdersByMonth")

	// Parse query parameters
	companyIdStr := r.URL.Query().Get("companyId")
	yearStr := r.URL.Query().Get("year")
	monthStr := r.URL.Query().Get("month")

	if companyIdStr == "" {
		json.NewEncoder(w).Encode(&types.Error{
			Status: http.StatusBadRequest,
			Error:  "Company ID is required",
		})
		return
	}

	if yearStr == "" {
		json.NewEncoder(w).Encode(&types.Error{
			Status: http.StatusBadRequest,
			Error:  "Year is required",
		})
		return
	}

	if monthStr == "" {
		json.NewEncoder(w).Encode(&types.Error{
			Status: http.StatusBadRequest,
			Error:  "Month is required",
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

	year, err := strconv.ParseInt(yearStr, 10, 32)
	if err != nil {
		json.NewEncoder(w).Encode(&types.Error{
			Status: http.StatusBadRequest,
			Error:  "Invalid year",
		})
		return
	}

	month, err := strconv.ParseInt(monthStr, 10, 32)
	if err != nil || month < 1 || month > 12 {
		json.NewEncoder(w).Encode(&types.Error{
			Status: http.StatusBadRequest,
			Error:  "Invalid month (must be 1-12)",
		})
		return
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
		fmt.Println("API Gateway :  GetOrdersByMonth - ERROR")
		errCompany.Error = "Company not found"
		json.NewEncoder(w).Encode(errCompany)
		return
	}

	res, errResp := orderClient.GetOrdersByMonth(companyId, int32(year), int32(month), r.Context())
	if errResp != nil {
		fmt.Println("API Gateway :  GetOrdersByMonth - ERROR")
		json.NewEncoder(w).Encode(errResp)
		return
	}

	fmt.Println("API Gateway :  GetOrdersByMonth - SUCCESS")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}
