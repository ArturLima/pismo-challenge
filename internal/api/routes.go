package api

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

func (api *Api) BindRoutes() {
	api.Router.Use(middleware.RequestID, middleware.Recoverer, middleware.Logger)

	api.Router.Post("/accounts", api.CreateAccount)
	api.Router.Get("/accounts/{accountId}", func(w http.ResponseWriter, r *http.Request) {})
	api.Router.Post("/transactions", func(w http.ResponseWriter, r *http.Request) {})
}
