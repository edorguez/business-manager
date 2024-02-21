package client

import (
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/gateway/pkg/client/routes"
	"github.com/EdoRguez/business-manager/gateway/pkg/config"
	"github.com/gorilla/mux"
)

func LoadRoutes(router *mux.Router, c *config.Config) {
	baseRoute := router.PathPrefix("/client").Subrouter()

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	getRouter := baseRoute.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {})

	postRouter := baseRoute.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("", svc.CreateClient)
}

func (svc *ServiceClient) CreateClient(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  CreateClient called --> 1")
	routes.CreateClient(w, r, svc.Client)
}
