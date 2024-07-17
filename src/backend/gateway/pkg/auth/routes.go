package auth

import (
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/gateway/pkg/auth/routes"
	"github.com/EdoRguez/business-manager/gateway/pkg/config"
	"github.com/gorilla/mux"
)

type AuthRoutes struct {
	config *config.Config
}

func LoadRoutes(router *mux.Router, c *config.Config) {
	loadUserRoutes(router, c)
	// loadPaymentRoutes(router, c)
}

func loadUserRoutes(router *mux.Router, c *config.Config) {
	baseRoute := router.PathPrefix("/users").Subrouter()

	ar := &AuthRoutes{
		config: c,
	}

	mw := MiddlewareConfig{}

	getRouter := baseRoute.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/{id:[0-9]+}", ar.GetUser)
	getRouter.HandleFunc("", ar.GetUsers)

	postRouter := baseRoute.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("", ar.CreateUser)
	postRouter.Use(mw.MiddlewareValidateCreateUser)

	putRouter := baseRoute.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ar.UpdateUser)
	putRouter.Use(mw.MiddlewareValidateUpdateUser)

	deleteRouter := baseRoute.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/{id:[0-9]+}", ar.DeleteUser)
}

// func loadPaymentRoutes(router *mux.Router, c *config.Config) {
// 	baseRoute := router.PathPrefix("/payments").Subrouter()
//
// 	cr := &CompanyRoutes{
// 		config: c,
// 	}
//
// 	mw := MiddlewareConfig{}
//
// 	getRouter := baseRoute.Methods(http.MethodGet).Subrouter()
// 	getRouter.HandleFunc("/{id:[0-9]+}", cr.GetPayment)
// 	getRouter.HandleFunc("", cr.GetPayments)
//
// 	postRouter := baseRoute.Methods(http.MethodPost).Subrouter()
// 	postRouter.HandleFunc("", cr.CreatePayment)
// 	postRouter.Use(mw.MiddlewareValidateCreatePayment)
//
// 	putRouter := baseRoute.Methods(http.MethodPut).Subrouter()
// 	putRouter.HandleFunc("/{id:[0-9]+}", cr.UpdatePayment)
// 	putRouter.Use(mw.MiddlewareValidateUpdatePayment)
//
// 	deleteRouter := baseRoute.Methods(http.MethodDelete).Subrouter()
// 	deleteRouter.HandleFunc("/{id:[0-9]+}", cr.DeletePayment)
// }

func (ar *AuthRoutes) CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  CreateUser Called --> 1")
	routes.CreateUser(w, r, ar.config)
}

func (ar *AuthRoutes) GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  GetUser Called --> 1")
	routes.GetUser(w, r, ar.config)
}

func (ar *AuthRoutes) GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  GetUsers Called --> 1")
	routes.GetUsers(w, r, ar.config)
}

func (ar *AuthRoutes) UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  UpdateUser Called --> 1")
	routes.UpdateUser(w, r, ar.config)
}

func (ar *AuthRoutes) DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  DeleteUser Called --> 1")
	routes.DeleteUser(w, r, ar.config)
}

// func (cr *CompanyRoutes) CreatePayment(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("API Gateway :  CreatePayment Called --> 1")
// 	routes.CreatePayment(w, r, cr.config)
// }
//
// func (cr *CompanyRoutes) GetPayment(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("API Gateway :  GetPayment Called --> 1")
// 	routes.GetPayment(w, r, cr.config)
// }
//
// func (cr *CompanyRoutes) GetPayments(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("API Gateway :  GetPayments Called --> 1")
// 	routes.GetPayments(w, r, cr.config)
// }
//
// func (cr *CompanyRoutes) UpdatePayment(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("API Gateway :  UpdatePayment Called --> 1")
// 	routes.UpdatePayment(w, r, cr.config)
// }
//
// func (cr *CompanyRoutes) DeletePayment(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("API Gateway :  DeletePayment Called --> 1")
// 	routes.DeletePayment(w, r, cr.config)
// }
