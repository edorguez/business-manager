package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/gateway/pkg/auth/client"
	"github.com/EdoRguez/business-manager/gateway/pkg/auth/contracts"
	"github.com/EdoRguez/business-manager/gateway/pkg/config"
)

type MiddlewareConfig struct {
	config *config.Config
}

func InitAuthMiddleware(c *config.Config) MiddlewareConfig {
	return MiddlewareConfig{config: c}
}

func (m *MiddlewareConfig) MiddlewareValidateAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			fmt.Println("API Gateway :  Middleware - Error - Validate Auth")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(&contracts.Error{
				Status: http.StatusUnauthorized,
				Error:  "Missing authorization header",
			})
			return
		}
		tokenString = tokenString[len("Bearer "):]

		if err := client.InitAuthServiceClient(m.config); err != nil {
			json.NewEncoder(w).Encode(&contracts.Error{
				Status: http.StatusInternalServerError,
				Error:  err.Error(),
			})
			return
		}

		validateToken, err := client.Validate(contracts.ValidateRequest{Token: tokenString}, r.Context())

		if err != nil {
			fmt.Println("API Gateway :  Validate - ERROR")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err)
			return
		}

		if validateToken.Status != http.StatusOK {
			fmt.Println("API Gateway :  Validate - ERROR")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(int(validateToken.Status))
			json.NewEncoder(w).Encode(&contracts.Error{
				Status: validateToken.Status,
				Error:  validateToken.Error,
			})
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (m *MiddlewareConfig) MiddlewareValidateLogin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body contracts.LoginRequest

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err := body.Validate()
		if err != nil {
			fmt.Println("API Gateway :  Middleware - Error - Login")
			middleErr := contracts.Error{
				Status: http.StatusBadRequest,
				Error:  err.Error(),
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(int(middleErr.Status))
			json.NewEncoder(w).Encode(middleErr)
			return
		}

		ctx := context.WithValue(r.Context(), contracts.LoginRequest{}, body)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
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
			middleErr := contracts.Error{
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
			middleErr := contracts.Error{
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
