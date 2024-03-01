package client

import (
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/gateway/pkg/client/routes"
	"github.com/EdoRguez/business-manager/gateway/pkg/config"
	"github.com/gorilla/mux"
)

func LoadRoutes(router *mux.Router, c *config.Config) {
	baseRoute := router.PathPrefix("/clients").Subrouter()

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	mw := MiddlewareConfig{}

	getRouter := baseRoute.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/{id:[0-9]+}", svc.GetClient)
	getRouter.HandleFunc("", svc.GetClients)

	postRouter := baseRoute.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("", svc.CreateClient)
	postRouter.Use(mw.MiddlewareValidateCreateClient)

	putRouter := baseRoute.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", svc.UpdateClient)
	putRouter.Use(mw.MiddlewareValidateUpdateClient)

	deleteRouter := baseRoute.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/{id:[0-9]+}", svc.DeleteClient)
}

func (svc *ServiceClient) CreateClient(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  CreateClient Called --> 1")
	routes.CreateClient(w, r, svc.Client)
}

func (svc *ServiceClient) GetClient(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  GetClient Called --> 1")
	routes.GetClient(w, r, svc.Client)
}

func (svc *ServiceClient) GetClients(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  GetClients Called --> 1")
	routes.GetClients(w, r, svc.Client)
}

func (svc *ServiceClient) UpdateClient(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  UpdateClient Called --> 1")
	routes.UpdateClient(w, r, svc.Client)
}

func (svc *ServiceClient) DeleteClient(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  DeleteClient Called --> 1")
	routes.DeleteClient(w, r, svc.Client)
}
