package customer

import (
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/gateway/pkg/config"
	"github.com/EdoRguez/business-manager/gateway/pkg/customer/routes"
	"github.com/gorilla/mux"
)

func LoadRoutes(router *mux.Router, c *config.Config) {
	baseRoute := router.PathPrefix("/customers").Subrouter()

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	mw := MiddlewareConfig{}

	getRouter := baseRoute.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/{id:[0-9]+}", svc.GetCustomer)
	getRouter.HandleFunc("", svc.GetCustomers)

	postRouter := baseRoute.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("", svc.CreateCustomer)
	postRouter.Use(mw.MiddlewareValidateCreateCustomer)

	putRouter := baseRoute.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", svc.UpdateCustomer)
	putRouter.Use(mw.MiddlewareValidateUpdateCustomer)

	deleteRouter := baseRoute.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/{id:[0-9]+}", svc.DeleteCustomer)
}

func (svc *ServiceClient) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  CreateCustomer Called --> 1")
	routes.CreateCustomer(w, r, svc.Client)
}

func (svc *ServiceClient) GetCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  GetCustomer Called --> 1")
	routes.GetCustomer(w, r, svc.Client)
}

func (svc *ServiceClient) GetCustomers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  GetCustomers Called --> 1")
	routes.GetCustomers(w, r, svc.Client)
}

func (svc *ServiceClient) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  UpdateCustomer Called --> 1")
	routes.UpdateCustomer(w, r, svc.Client)
}

func (svc *ServiceClient) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  DeleteCustomer Called --> 1")
	routes.DeleteCustomer(w, r, svc.Client)
}
