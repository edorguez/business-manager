package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/gateway/pkg/auth/contracts"
)

type MiddlewareConfig struct{}

type MiddlewareErrorResponse struct {
	Status int64
	Error  string
}

func (m *MiddlewareConfig) MiddlewareValidateCreateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body contracts.CreateUserRequest

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err := body.Validate()
		if err != nil {
			fmt.Println("API Gateway :  Middleware - Error - CreateUser")
			middleErr := MiddlewareErrorResponse{
				Status: http.StatusBadRequest,
				Error:  err.Error(),
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(int(middleErr.Status))
			json.NewEncoder(w).Encode(middleErr)
			return
		}

		ctx := context.WithValue(r.Context(), contracts.CreateUserRequest{}, body)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}

func (m *MiddlewareConfig) MiddlewareValidateUpdateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body contracts.UpdateUserRequest

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err := body.Validate()
		if err != nil {
			fmt.Println("API Gateway :  Middleware - Error - UpdateUser")
			middleErr := MiddlewareErrorResponse{
				Status: http.StatusBadRequest,
				Error:  err.Error(),
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(int(middleErr.Status))
			json.NewEncoder(w).Encode(middleErr)
			return
		}

		ctx := context.WithValue(r.Context(), contracts.UpdateUserRequest{}, body)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
