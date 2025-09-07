package company

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/edorguez/business-manager/services/gateway/pkg/company/contracts"
)

type MiddlewareConfig struct{}

func (m *MiddlewareConfig) MiddlewareValidateCreateCompany(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body contracts.CreateCompanyRequest

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err := body.Validate()
		if err != nil {
			fmt.Println("API Gateway :  Middleware - Error - CreateCompany")
			middleErr := contracts.Error{
				Status: http.StatusBadRequest,
				Error:  err.Error(),
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(int(middleErr.Status))
			json.NewEncoder(w).Encode(middleErr)
			return
		}

		ctx := context.WithValue(r.Context(), contracts.CreateCompanyRequest{}, body)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}

func (m *MiddlewareConfig) MiddlewareValidateUpdateCompany(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body contracts.UpdateCompanyRequest

		// Parse the form data
		err := r.ParseMultipartForm(10 << 20) // 10 MB max memory
		if err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}

		// Get the JSON part
		jsonData := r.FormValue("json")
		err = json.Unmarshal([]byte(jsonData), &body)
		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = body.Validate()
		if err != nil {
			fmt.Println("API Gateway :  Middleware - Error - UpdateCompany")
			middleErr := contracts.Error{
				Status: http.StatusBadRequest,
				Error:  err.Error(),
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(int(middleErr.Status))
			json.NewEncoder(w).Encode(middleErr)
			return
		}

		ctx := context.WithValue(r.Context(), contracts.UpdateCompanyRequest{}, body)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}

func (m *MiddlewareConfig) MiddlewareValidateCreatePayment(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body contracts.CreatePaymentRequest

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err := body.Validate()
		if err != nil {
			fmt.Println("API Gateway :  Middleware - Error - CreatePayment")
			middleErr := contracts.Error{
				Status: http.StatusBadRequest,
				Error:  err.Error(),
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(int(middleErr.Status))
			json.NewEncoder(w).Encode(middleErr)
			return
		}

		ctx := context.WithValue(r.Context(), contracts.CreatePaymentRequest{}, body)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}

func (m *MiddlewareConfig) MiddlewareValidateUpdatePayment(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body contracts.UpdatePaymentRequest

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err := body.Validate()
		if err != nil {
			fmt.Println("API Gateway :  Middleware - Error - UpdatePayment")
			middleErr := contracts.Error{
				Status: http.StatusBadRequest,
				Error:  err.Error(),
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(int(middleErr.Status))
			json.NewEncoder(w).Encode(middleErr)
			return
		}

		ctx := context.WithValue(r.Context(), contracts.UpdatePaymentRequest{}, body)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}

func (m *MiddlewareConfig) MiddlewareValidateUpdatePaymentStatus(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body contracts.UpdatePaymentStatusRequest

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err := body.Validate()
		if err != nil {
			fmt.Println("API Gateway :  Middleware - Error - UpdatePaymentStatus")
			middleErr := contracts.Error{
				Status: http.StatusBadRequest,
				Error:  err.Error(),
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(int(middleErr.Status))
			json.NewEncoder(w).Encode(middleErr)
			return
		}

		ctx := context.WithValue(r.Context(), contracts.UpdatePaymentStatusRequest{}, body)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
