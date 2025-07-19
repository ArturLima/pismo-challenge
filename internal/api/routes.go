package api

import (
	"github.com/go-chi/chi/v5/middleware"
)

func (api *Api) BindRoutes() {
	api.Router.Use(middleware.RequestID, middleware.Recoverer, middleware.Logger)

	api.Router.Post("/accounts", api.CreateAccount)
	api.Router.Get("/accounts/{accountId}", api.GetAccount)
	api.Router.Post("/transactions", api.CreateTransaction)
}
