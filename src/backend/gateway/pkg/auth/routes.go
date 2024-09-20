package auth

import (
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/gateway/pkg/auth/routes"
	"github.com/EdoRguez/business-manager/gateway/pkg/config"
	"github.com/gorilla/mux"
)

type AuthRoutes struct {
	config *config.Config
}

func LoadRoutes(router *mux.Router, c *config.Config) {
	loadUserRoutes(router, c)
	loadAuthRoutes(router, c)
}

func loadUserRoutes(router *mux.Router, c *config.Config) {
	baseRoute := router.PathPrefix("/users").Subrouter()

	ar := &AuthRoutes{
		config: c,
	}

	mw := InitAuthMiddleware(c)
	baseRoute.Use(mw.MiddlewareValidateAuth)

	getRouter := baseRoute.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/{id:[0-9]+}", ar.GetUser)
	getRouter.HandleFunc("", ar.GetUsers)

	postRouter := baseRoute.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("", ar.CreateUser)
	postRouter.Use(mw.MiddlewareValidateCreateUser)

	putRouteEmail := baseRoute.Methods(http.MethodPut).Subrouter()
	putRouteEmail.HandleFunc("/{id:[0-9]+}/email", ar.UpdateEmail)
	putRouteEmail.Use(mw.MiddlewareValidateUpdateEmail)

	putRoutePassword := baseRoute.Methods(http.MethodPut).Subrouter()
	putRoutePassword.HandleFunc("/{id:[0-9]+}/password", ar.UpdatePassword)
	putRoutePassword.Use(mw.MiddlewareValidateUpdatePassword)

	putRouter := baseRoute.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ar.UpdateUser)
	putRouter.Use(mw.MiddlewareValidateUpdateUser)

	deleteRouter := baseRoute.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/{id:[0-9]+}", ar.DeleteUser)
}

func loadAuthRoutes(router *mux.Router, c *config.Config) {
	baseRoute := router.PathPrefix("/auth").Subrouter()

	ar := &AuthRoutes{
		config: c,
	}

	mw := InitAuthMiddleware(c)

	registerRouter := baseRoute.Methods(http.MethodPost).Subrouter()
	registerRouter.HandleFunc("/register", ar.Register)
	// registerRouter.Use(mw.MiddlewareValidateAuth)
	registerRouter.Use(mw.MiddlewareValidateCreateUser)

	loginRouter := baseRoute.Methods(http.MethodPost).Subrouter()
	loginRouter.HandleFunc("/login", ar.Login)
	loginRouter.Use(mw.MiddlewareValidateLogin)
}

func (ar *AuthRoutes) CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  CreateUser Called --> 1")
	routes.CreateUser(w, r, ar.config)
}

func (ar *AuthRoutes) GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  GetUser Called --> 1")
	routes.GetUser(w, r, ar.config)
}

func (ar *AuthRoutes) GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  GetUsers Called --> 1")
	routes.GetUsers(w, r, ar.config)
}

func (ar *AuthRoutes) UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  UpdateUser Called --> 1")
	routes.UpdateUser(w, r, ar.config)
}

func (ar *AuthRoutes) UpdateEmail(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  UpdateEmail Called --> 1")
	routes.UpdateEmail(w, r, ar.config)
}

func (ar *AuthRoutes) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  UpdatePassword Called --> 1")
	routes.UpdatePassword(w, r, ar.config)
}

func (ar *AuthRoutes) DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  DeleteUser Called --> 1")
	routes.DeleteUser(w, r, ar.config)
}

func (ar *AuthRoutes) Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  Register Called --> 1")
	routes.Register(w, r, ar.config)
}

func (ar *AuthRoutes) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  Login Called --> 1")
	routes.Login(w, r, ar.config)
}
