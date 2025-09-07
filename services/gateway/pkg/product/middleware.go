package product

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/edorguez/business-manager/services/gateway/pkg/product/contracts"
)

type MiddlewareConfig struct{}

func (m *MiddlewareConfig) MiddlewareValidateCreateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body contracts.CreateProductRequest

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
			fmt.Println("API Gateway :  Middleware - Error - CreateProduct")
			fmt.Println(err)
			middleErr := contracts.Error{
				Status: http.StatusBadRequest,
				Error:  err.Error(),
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(int(middleErr.Status))
			json.NewEncoder(w).Encode(middleErr)
			return
		}

		ctx := context.WithValue(r.Context(), "keyProductCreate", body)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}

func (m *MiddlewareConfig) MiddlewareValidateUpdateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body contracts.UpdateProductRequest

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
			fmt.Println("API Gateway :  Middleware - Error - UpdateProduct")
			middleErr := contracts.Error{
				Status: http.StatusBadRequest,
				Error:  err.Error(),
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(int(middleErr.Status))
			json.NewEncoder(w).Encode(middleErr)
			return
		}

		ctx := context.WithValue(r.Context(), "keyProductUpdate", body)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}

func (m *MiddlewareConfig) MiddlewareValidateUpdateProductStatus(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body contracts.UpdateProductStatusRequest

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err := body.Validate()
		if err != nil {
			fmt.Println("API Gateway :  Middleware - Error - UpdateProductStatus")
			middleErr := contracts.Error{
				Status: http.StatusBadRequest,
				Error:  err.Error(),
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(int(middleErr.Status))
			json.NewEncoder(w).Encode(middleErr)
			return
		}

		ctx := context.WithValue(r.Context(), "keyProductStatusUpdate", body)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
