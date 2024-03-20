package company

import (
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/gateway/pkg/company/routes"
	"github.com/EdoRguez/business-manager/gateway/pkg/config"
	"github.com/gorilla/mux"
)

func LoadRoutes(router *mux.Router, c *config.Config) {
	baseRoute := router.PathPrefix("/companies").Subrouter()

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	mw := MiddlewareConfig{}

	getRouter := baseRoute.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/{id:[0-9]+}", svc.GetCompany)
	getRouter.HandleFunc("", svc.GetCompanies)

	postRouter := baseRoute.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("", svc.CreateCompany)
	postRouter.Use(mw.MiddlewareValidateCreateCompany)

	putRouter := baseRoute.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", svc.UpdateCompany)
	putRouter.Use(mw.MiddlewareValidateUpdateCompany)

	deleteRouter := baseRoute.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/{id:[0-9]+}", svc.DeleteCompany)
}

func (svc *ServiceClient) CreateCompany(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  CreateCompany Called --> 1")
	routes.CreateCompany(w, r, svc.Client)
}

func (svc *ServiceClient) GetCompany(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  GetCompany Called --> 1")
	routes.GetCompany(w, r, svc.Client)
}

func (svc *ServiceClient) GetCompanies(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  GetCompanies Called --> 1")
	routes.GetCompanies(w, r, svc.Client)
}

func (svc *ServiceClient) UpdateCompany(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  UpdateCompany Called --> 1")
	routes.UpdateCompany(w, r, svc.Client)
}

func (svc *ServiceClient) DeleteCompany(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  DeleteCompany Called --> 1")
	routes.DeleteCompany(w, r, svc.Client)
}
