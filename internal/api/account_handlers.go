package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (api *Api) CreateAccount(w http.ResponseWriter, r *http.Request) {
	// Handler logic for creating an account
}

func (api *Api) GetAccount(w http.ResponseWriter, r *http.Request) {
	accountId := chi.URLParam(r, "accountId")
	fmt.Println("Getting account with ID:", accountId)
	// Handler logic for getting an account
}
