package product

import (
	"fmt"
	"net/http"

	auth "github.com/EdoRguez/business-manager/gateway/pkg/auth"
	"github.com/EdoRguez/business-manager/gateway/pkg/config"
	"github.com/EdoRguez/business-manager/gateway/pkg/product/routes"
	"github.com/gorilla/mux"
)

type ProductRoutes struct {
	config *config.Config
}

func LoadRoutes(router *mux.Router, c *config.Config) {
	baseRoute := router.PathPrefix("/products").Subrouter()

	mwc := auth.InitAuthMiddleware(c)

	cr := &ProductRoutes{
		config: c,
	}

	mw := MiddlewareConfig{}

	getRouter := baseRoute.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/latest", cr.GetLatestProducts)
	getRouter.HandleFunc("/{id}", cr.GetProduct)
	getRouter.Use(mwc.MiddlewareValidateAuth)

	getProductsRouter := baseRoute.Methods(http.MethodGet).Subrouter()
	getProductsRouter.HandleFunc("", cr.GetProducts)

	postRouter := baseRoute.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("", cr.CreateProduct)
	postRouter.Use(mw.MiddlewareValidateCreateProduct)
	postRouter.Use(mwc.MiddlewareValidateAuth)

	putRouter := baseRoute.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id}", cr.UpdateProduct)
	putRouter.Use(mw.MiddlewareValidateUpdateProduct)
	putRouter.Use(mwc.MiddlewareValidateAuth)

	putStatusRoutes := baseRoute.Methods(http.MethodPut).Subrouter()
	putStatusRoutes.HandleFunc("/{id}/status", cr.UpdateProductStatus)
	putStatusRoutes.Use(mwc.MiddlewareValidateAuth)
	putStatusRoutes.Use(mw.MiddlewareValidateUpdateProductStatus)

	deleteRouter := baseRoute.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/{id}", cr.DeleteProduct)
	deleteRouter.Use(mwc.MiddlewareValidateAuth)
}

func (pr *ProductRoutes) CreateProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  CreateProduct Called --> 1")
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

func (cr *ProductRoutes) GetLatestProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  GetLatestProducts Called --> 1")
	routes.GetLatestProducts(w, r, cr.config)
}

func (cr *ProductRoutes) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  UpdateProduct Called --> 1")
	routes.UpdateProduct(w, r, cr.config)
}

func (cr *ProductRoutes) UpdateProductStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  UpdateProductStatus Called --> 1")
	routes.UpdateProductStatus(w, r, cr.config)
}

func (cr *ProductRoutes) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  DeleteProduct Called --> 1")
	routes.DeleteProduct(w, r, cr.config)
}
