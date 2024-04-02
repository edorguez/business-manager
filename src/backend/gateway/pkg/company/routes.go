package company

import (
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/gateway/pkg/company/routes"
	"github.com/EdoRguez/business-manager/gateway/pkg/config"
	"github.com/gorilla/mux"
)

func LoadRoutes(router *mux.Router, c *config.Config) {
	svc := InitServiceClient(c)

	loadCompanyRoutes(router, svc)
	loadPaymentRoutes(router, svc)
}

func loadCompanyRoutes(router *mux.Router, svc ServiceClient) {
	baseRoute := router.PathPrefix("/companies").Subrouter()

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

func loadPaymentRoutes(router *mux.Router, svc ServiceClient) {
	baseRoute := router.PathPrefix("/payments").Subrouter()

	mw := MiddlewareConfig{}

	getRouter := baseRoute.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/{id:[0-9]+}", svc.GetPayment)
	getRouter.HandleFunc("", svc.GetPayments)

	postRouter := baseRoute.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("", svc.CreatePayment)
	postRouter.Use(mw.MiddlewareValidateCreatePayment)

	putRouter := baseRoute.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", svc.UpdatePayment)
	putRouter.Use(mw.MiddlewareValidateUpdatePayment)

	deleteRouter := baseRoute.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/{id:[0-9]+}", svc.DeletePayment)
}

func (svc *ServiceClient) CreateCompany(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  CreateCompany Called --> 1")
	routes.CreateCompany(w, r, svc.CompanyClient)
}

func (svc *ServiceClient) GetCompany(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  GetCompany Called --> 1")
	routes.GetCompany(w, r, svc.CompanyClient)
}

func (svc *ServiceClient) GetCompanies(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  GetCompanies Called --> 1")
	routes.GetCompanies(w, r, svc.CompanyClient)
}

func (svc *ServiceClient) UpdateCompany(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  UpdateCompany Called --> 1")
	routes.UpdateCompany(w, r, svc.CompanyClient)
}

func (svc *ServiceClient) DeleteCompany(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  DeleteCompany Called --> 1")
	routes.DeleteCompany(w, r, svc.CompanyClient)
}

func (svc *ServiceClient) CreatePayment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  CreatePayment Called --> 1")
	routes.CreatePayment(w, r, svc.PaymentClient)
}

func (svc *ServiceClient) GetPayment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  GetPayment Called --> 1")
	routes.GetPayment(w, r, svc.PaymentClient)
}

func (svc *ServiceClient) GetPayments(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  GetPayments Called --> 1")
	routes.GetPayments(w, r, svc.PaymentClient)
}

func (svc *ServiceClient) UpdatePayment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  UpdatePayment Called --> 1")
	routes.UpdatePayment(w, r, svc.PaymentClient)
}

func (svc *ServiceClient) DeletePayment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  DeletePayment Called --> 1")
	routes.DeletePayment(w, r, svc.PaymentClient)
}
