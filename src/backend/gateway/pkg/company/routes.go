package company

import (
	"fmt"
	"net/http"

	auth "github.com/EdoRguez/business-manager/gateway/pkg/auth"
	"github.com/EdoRguez/business-manager/gateway/pkg/company/routes"
	"github.com/EdoRguez/business-manager/gateway/pkg/config"
	"github.com/gorilla/mux"
)

type CompanyRoutes struct {
	config *config.Config
}

func LoadRoutes(router *mux.Router, c *config.Config) {
	loadCompanyRoutes(router, c)
	loadPaymentRoutes(router, c)
	loadPaymentTypeRoutes(router, c)
}

func loadCompanyRoutes(router *mux.Router, c *config.Config) {
	baseRoute := router.PathPrefix("/companies").Subrouter()

	mwc := auth.InitAuthMiddleware(c)

	cr := &CompanyRoutes{
		config: c,
	}

	mw := MiddlewareConfig{}

	getRouter := baseRoute.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/{id:[0-9]+}", cr.GetCompany)
	getRouter.HandleFunc("", cr.GetCompanies)
	getRouter.Use(mwc.MiddlewareValidateAuth)

	getCompanyByNameRouter := baseRoute.Methods(http.MethodGet).Subrouter()
	getCompanyByNameRouter.HandleFunc("/name/{name}", cr.GetCompanyByName)
	getCompanyByNameRouter.HandleFunc("/nameUrl/{nameUrl}", cr.GetCompanyByNameUrl)

	postRouter := baseRoute.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("", cr.CreateCompany)
	postRouter.Use(mw.MiddlewareValidateCreateCompany)
	postRouter.Use(mwc.MiddlewareValidateAuth)

	putRouter := baseRoute.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", cr.UpdateCompany)
	putRouter.Use(mw.MiddlewareValidateUpdateCompany)
	putRouter.Use(mwc.MiddlewareValidateAuth)

	deleteRouter := baseRoute.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/{id:[0-9]+}", cr.DeleteCompany)
	deleteRouter.Use(mwc.MiddlewareValidateAuth)
}

func loadPaymentRoutes(router *mux.Router, c *config.Config) {
	baseRoute := router.PathPrefix("/payments").Subrouter()

	mwc := auth.InitAuthMiddleware(c)
	baseRoute.Use(mwc.MiddlewareValidateAuth)

	cr := &CompanyRoutes{
		config: c,
	}

	mw := MiddlewareConfig{}

	getRouter := baseRoute.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/types", cr.GetPaymentsTypes)
	getRouter.HandleFunc("/{id:[0-9]+}", cr.GetPayment)
	getRouter.HandleFunc("", cr.GetPayments)

	postRouter := baseRoute.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("", cr.CreatePayment)
	postRouter.Use(mw.MiddlewareValidateCreatePayment)

	putRouter := baseRoute.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", cr.UpdatePayment)
	putRouter.Use(mw.MiddlewareValidateUpdatePayment)

	putStatusRoutes := baseRoute.Methods(http.MethodPut).Subrouter()
	putStatusRoutes.HandleFunc("/{id:[0-9]+}/status", cr.UpdatePaymentStatus)
	// putStatusRoutes.Use(mw.MiddlewareValidateUpdatePaymentStatus)

	deleteRouter := baseRoute.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/{id:[0-9]+}", cr.DeletePayment)
}

func loadPaymentTypeRoutes(router *mux.Router, c *config.Config) {
	baseRoute := router.PathPrefix("/paymentTypes").Subrouter()

	mwc := auth.InitAuthMiddleware(c)
	baseRoute.Use(mwc.MiddlewareValidateAuth)

	cr := &CompanyRoutes{
		config: c,
	}

	getRouter := baseRoute.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/{id:[0-9]+}", cr.GetPaymentType)
	getRouter.HandleFunc("", cr.GetPaymentTypes)
}

func (cr *CompanyRoutes) CreateCompany(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  CreateCompany Called --> 1")
	routes.CreateCompany(w, r, cr.config)
}

func (cr *CompanyRoutes) GetCompany(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  GetCompany Called --> 1")
	routes.GetCompany(w, r, cr.config)
}

func (cr *CompanyRoutes) GetCompanyByName(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  GetCompanyByName Called --> 1")
	routes.GetCompanyByName(w, r, cr.config)
}

func (cr *CompanyRoutes) GetCompanyByNameUrl(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  GetCompanyByNameUrl Called --> 1")
	routes.GetCompanyByNameUrl(w, r, cr.config)
}

func (cr *CompanyRoutes) GetCompanies(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  GetCompanies Called --> 1")
	routes.GetCompanies(w, r, cr.config)
}

func (cr *CompanyRoutes) UpdateCompany(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  UpdateCompany Called --> 1")
	routes.UpdateCompany(w, r, cr.config)
}

func (cr *CompanyRoutes) DeleteCompany(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  DeleteCompany Called --> 1")
	routes.DeleteCompany(w, r, cr.config)
}

func (cr *CompanyRoutes) CreatePayment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  CreatePayment Called --> 1")
	routes.CreatePayment(w, r, cr.config)
}

func (cr *CompanyRoutes) GetPayment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  GetPayment Called --> 1")
	routes.GetPayment(w, r, cr.config)
}

func (cr *CompanyRoutes) GetPayments(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  GetPayments Called --> 1")
	routes.GetPayments(w, r, cr.config)
}

func (cr *CompanyRoutes) GetPaymentsTypes(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  GetPaymentsTypes Called --> 1")
	routes.GetPaymentsTypes(w, r, cr.config)
}

func (cr *CompanyRoutes) UpdatePayment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  UpdatePayment Called --> 1")
	routes.UpdatePayment(w, r, cr.config)
}

func (cr *CompanyRoutes) UpdatePaymentStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  UpdatePayment Called --> 1")
	routes.UpdatePaymentStatus(w, r, cr.config)
}

func (cr *CompanyRoutes) DeletePayment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  DeletePayment Called --> 1")
	routes.DeletePayment(w, r, cr.config)
}

func (cr *CompanyRoutes) GetPaymentType(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  GetPaymentType Called --> 1")
	routes.GetPaymentType(w, r, cr.config)
}

func (cr *CompanyRoutes) GetPaymentTypes(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  GetPaymentTypes Called --> 1")
	routes.GetPaymentTypes(w, r, cr.config)
}
