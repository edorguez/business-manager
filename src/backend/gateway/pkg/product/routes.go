package product

import (
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/gateway/pkg/config"
	"github.com/EdoRguez/business-manager/gateway/pkg/product/routes"
	"github.com/gorilla/mux"
)

type ProductRoutes struct {
	config *config.Config
}

func LoadRoutes(router *mux.Router, c *config.Config) {
	baseRoute := router.PathPrefix("/products").Subrouter()

	cr := &ProductRoutes{
		config: c,
	}

	mw := MiddlewareConfig{}

	// getRouter := baseRoute.Methods(http.MethodGet).Subrouter()
	// getRouter.HandleFunc("/{id:[0-9]+}", cr.GetCustomer)
	// getRouter.HandleFunc("", cr.GetCustomers)

	postRouter := baseRoute.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("", cr.CreateProduct)
	postRouter.Use(mw.MiddlewareValidateCreateProduct)

	// putRouter := baseRoute.Methods(http.MethodPut).Subrouter()
	// putRouter.HandleFunc("/{id:[0-9]+}", cr.UpdateCustomer)
	// putRouter.Use(mw.MiddlewareValidateUpdateCustomer)
	//
	// deleteRouter := baseRoute.Methods(http.MethodDelete).Subrouter()
	// deleteRouter.HandleFunc("/{id:[0-9]+}", cr.DeleteCustomer)
}

func (pr *ProductRoutes) CreateProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  Createproduct Called --> 1")
	routes.CreateProduct(w, r, pr.config)
}

// func (cr *CustomerRoutes) GetCustomer(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("API Gateway :  GetCustomer Called --> 1")
// 	routes.GetCustomer(w, r, cr.config)
// }
//
// func (cr *CustomerRoutes) GetCustomers(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("API Gateway :  GetCustomers Called --> 1")
// 	routes.GetCustomers(w, r, cr.config)
// }
//
// func (cr *CustomerRoutes) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("API Gateway :  UpdateCustomer Called --> 1")
// 	routes.UpdateCustomer(w, r, cr.config)
// }
//
// func (cr *CustomerRoutes) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("API Gateway :  DeleteCustomer Called --> 1")
// 	routes.DeleteCustomer(w, r, cr.config)
// }
