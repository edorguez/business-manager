package order

import (
	"fmt"
	"net/http"

	"github.com/edorguez/business-manager/services/gateway/pkg/config"
	"github.com/edorguez/business-manager/services/gateway/pkg/order/routes"
	"github.com/gorilla/mux"
)

type OrderRoutes struct {
	config *config.Config
}

func LoadRoutes(router *mux.Router, c *config.Config) {
	baseRoute := router.PathPrefix("/orders").Subrouter()

	cr := &OrderRoutes{
		config: c,
	}

	mw := MiddlewareConfig{}

	postRouter := baseRoute.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("", cr.CreateOrder)
	postRouter.Use(mw.MiddlewareValidateCreateOrder)
}

func (or *OrderRoutes) CreateOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  CreateOrder Called --> 1")
	routes.CreateOrder(w, r, or.config)
}
