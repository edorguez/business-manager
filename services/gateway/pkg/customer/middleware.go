package customer

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/edorguez/business-manager/services/gateway/pkg/customer/contracts"
)

type MiddlewareConfig struct{}

func (m *MiddlewareConfig) MiddlewareValidateCreateCustomer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body contracts.CreateCustomerRequest

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err := body.Validate()
		if err != nil {
			fmt.Println("API Gateway :  Middleware - Error - CreateCustomer")
			middleErr := contracts.Error{
				Status: http.StatusBadRequest,
				Error:  err.Error(),
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(int(middleErr.Status))
			json.NewEncoder(w).Encode(middleErr)
			return
		}

		ctx := context.WithValue(r.Context(), contracts.CreateCustomerRequest{}, body)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}

func (m *MiddlewareConfig) MiddlewareValidateUpdateCustomer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body contracts.UpdateCustomerRequest

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err := body.Validate()
		if err != nil {
			fmt.Println("API Gateway :  Middleware - Error - UpdateCustomer")
			middleErr := contracts.Error{
				Status: http.StatusBadRequest,
				Error:  err.Error(),
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(int(middleErr.Status))
			json.NewEncoder(w).Encode(middleErr)
			return
		}

		ctx := context.WithValue(r.Context(), contracts.UpdateCustomerRequest{}, body)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
