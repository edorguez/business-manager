package whatsapp

import (
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/gateway/pkg/auth"
	"github.com/EdoRguez/business-manager/gateway/pkg/config"
	"github.com/EdoRguez/business-manager/gateway/pkg/whatsapp/routes"
	"github.com/gorilla/mux"
)

type WhatsappRoutes struct {
	config *config.Config
}

func LoadRoutes(router *mux.Router, c *config.Config) {
	baseRoute := router.PathPrefix("/whatsapp").Subrouter()

	mwc := auth.InitAuthMiddleware(c)
	baseRoute.Use(mwc.MiddlewareValidateAuth)

	wr := &WhatsappRoutes{
		config: c,
	}

	mw := MiddlewareConfig{}

	putRouter := baseRoute.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/businessNumber", wr.UpdateBusinessNumber)
	putRouter.Use(mw.MiddlewareValidateUpdateBusinessPhone)

}

func (wr *WhatsappRoutes) UpdateBusinessNumber(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway : UpdateBusinessNumber Called --> 1")
	routes.UpdateBusinessNumber(w, r, wr.config)
}
