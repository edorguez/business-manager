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

	getRouter := baseRoute.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/businessPhone/{id}", wr.GetBusinessPhoneByCompanyId)

	postRouter := baseRoute.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/businessPhone", wr.CreateBusinessPhone)
	postRouter.Use(mw.MiddlewareValidateCreateBusinessPhone)

	putRouter := baseRoute.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/businessPhone", wr.UpdateBusinessPhone)
	putRouter.Use(mw.MiddlewareValidateUpdateBusinessPhone)

}

func (wr *WhatsappRoutes) GetBusinessPhoneByCompanyId(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway : GetBusinessPhoneByCompanyId Called --> 1")
	routes.GetBusinessPhoneByCompanyId(w, r, wr.config)
}

func (wr *WhatsappRoutes) CreateBusinessPhone(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway : CreateBusinessPhone Called --> 1")
	routes.CreateBusinessPhone(w, r, wr.config)
}

func (wr *WhatsappRoutes) UpdateBusinessPhone(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway : UpdateBusinessPhone Called --> 1")
	routes.UpdateBusinessPhone(w, r, wr.config)
}
