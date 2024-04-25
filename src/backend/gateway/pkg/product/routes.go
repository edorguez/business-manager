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

	getRouter := baseRoute.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/{id}", cr.GetProduct)
	getRouter.HandleFunc("", cr.GetProducts)

	postRouter := baseRoute.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("", cr.CreateProduct)
	postRouter.Use(mw.MiddlewareValidateCreateProduct)

	putRouter := baseRoute.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id}", cr.UpdateProduct)
	putRouter.Use(mw.MiddlewareValidateUpdateProduct)

	deleteRouter := baseRoute.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/{id}", cr.DeleteProduct)
}

func (pr *ProductRoutes) CreateProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  Createproduct Called --> 1")
	routes.CreateProduct(w, r, pr.config)
}

func (cr *ProductRoutes) GetProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  GetProduct Called --> 1")
	routes.GetProduct(w, r, cr.config)
}

func (cr *ProductRoutes) GetProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  GetProducts Called --> 1")
	routes.GetProducts(w, r, cr.config)
}

func (cr *ProductRoutes) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  UpdateProduct Called --> 1")
	routes.UpdateProduct(w, r, cr.config)
}

func (cr *ProductRoutes) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  DeleteProduct Called --> 1")
	routes.DeleteProduct(w, r, cr.config)
}
