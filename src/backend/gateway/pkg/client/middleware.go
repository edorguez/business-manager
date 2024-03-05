package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/gateway/pkg/client/routes"
)

type MiddlewareConfig struct{}

type MiddlewareErrorResponse struct {
	Status int64
	Error  string
}

func (m *MiddlewareConfig) MiddlewareValidateCreateClient(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body routes.CreateClientRequestBody

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err := body.Validate()
		if err != nil {
			fmt.Println("API Gateway :  Middleware - Error - CreateClient")
			middleErr := MiddlewareErrorResponse{
				Status: http.StatusBadRequest,
				Error:  err.Error(),
			}
			w.WriteHeader(int(middleErr.Status))
			json.NewEncoder(w).Encode(middleErr)
			return
		}

		ctx := context.WithValue(r.Context(), routes.CreateClientRequestBody{}, body)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}

func (m *MiddlewareConfig) MiddlewareValidateUpdateClient(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body routes.UpdateClientRequestBody

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err := body.Validate()
		if err != nil {
			fmt.Println("API Gateway :  Middleware - Error - UpdateClient")
			middleErr := MiddlewareErrorResponse{
				Status: http.StatusBadRequest,
				Error:  err.Error(),
			}
			w.WriteHeader(int(middleErr.Status))
			json.NewEncoder(w).Encode(middleErr)
			return
		}

		ctx := context.WithValue(r.Context(), routes.UpdateClientRequestBody{}, body)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
