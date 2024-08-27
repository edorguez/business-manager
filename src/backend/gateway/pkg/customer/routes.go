package customer

import (
	"fmt"
	"net/http"

	auth "github.com/EdoRguez/business-manager/gateway/pkg/auth"
	"github.com/EdoRguez/business-manager/gateway/pkg/config"
	"github.com/EdoRguez/business-manager/gateway/pkg/customer/routes"
	"github.com/gorilla/mux"
)

type CustomerRoutes struct {
	config *config.Config
}

func LoadRoutes(router *mux.Router, c *config.Config) {
	baseRoute := router.PathPrefix("/customers").Subrouter()

	mwc := auth.InitAuthMiddleware(c)
	baseRoute.Use(mwc.MiddlewareValidateAuth)

	cr := &CustomerRoutes{
		config: c,
	}

	mw := MiddlewareConfig{}

	getRouter := baseRoute.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/{id:[0-9]+}", cr.GetCustomer)
	getRouter.HandleFunc("/months", cr.GetCustomersByMonths)
	getRouter.HandleFunc("", cr.GetCustomers)

	postRouter := baseRoute.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("", cr.CreateCustomer)
	postRouter.Use(mw.MiddlewareValidateCreateCustomer)

	putRouter := baseRoute.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", cr.UpdateCustomer)
	putRouter.Use(mw.MiddlewareValidateUpdateCustomer)

	deleteRouter := baseRoute.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/{id:[0-9]+}", cr.DeleteCustomer)
}

func (cr *CustomerRoutes) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  CreateCustomer Called --> 1")
	routes.CreateCustomer(w, r, cr.config)
}

func (cr *CustomerRoutes) GetCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  GetCustomer Called --> 1")
	routes.GetCustomer(w, r, cr.config)
}

func (cr *CustomerRoutes) GetCustomers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  GetCustomers Called --> 1")
	routes.GetCustomers(w, r, cr.config)
}

func (cr *CustomerRoutes) GetCustomersByMonths(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  GetCustomersByMonths Called --> 1")
	routes.GetCustomersByMonths(w, r, cr.config)
}

func (cr *CustomerRoutes) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  UpdateCustomer Called --> 1")
	routes.UpdateCustomer(w, r, cr.config)
}

func (cr *CustomerRoutes) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  DeleteCustomer Called --> 1")
	routes.DeleteCustomer(w, r, cr.config)
}
