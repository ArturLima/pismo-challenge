package api

import (
	"github.com/ArturLima/pismo/internal/services"
	"github.com/go-chi/chi/v5"
)

type Api struct {
	Router             *chi.Mux
	AccountService     services.IAccountService
	TransactionService services.ITransactionService
}
