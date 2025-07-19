package api

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/ArturLima/pismo/internal/useCases/account"
	"github.com/ArturLima/pismo/internal/utils"
	"github.com/go-chi/chi/v5"
)

func (api *Api) CreateAccount(w http.ResponseWriter, r *http.Request) {
	data, err := utils.DecodeJSON[account.CreateAccountRequest](r)
	if err != nil {
		log.Printf("Error decoding JSON: %v", err)
		utils.EncodeJSON(w, r, http.StatusBadRequest, map[string]any{"error": "Invalid request body"})
		return
	}

	account, err := api.AccountService.CreateAccount(r.Context(), data.Document)
	if err != nil {
		log.Printf("Error creating account: %v", err)
		utils.EncodeJSON(w, r, http.StatusInternalServerError, map[string]any{"error": "Failed to create account"})
		return
	}
	utils.EncodeJSON(w, r, http.StatusCreated, account)
}

func (api *Api) GetAccount(w http.ResponseWriter, r *http.Request) {
	accountId := chi.URLParam(r, "accountId")

	if strings.TrimSpace(accountId) == "" {
		utils.EncodeJSON(w, r, http.StatusBadRequest, map[string]any{"error": "Account ID is required"})
		return
	}
	id, err := strconv.Atoi(accountId)
	if err != nil {
		utils.EncodeJSON(w, r, http.StatusBadRequest, map[string]any{"error": "Invalid account ID format"})
		return
	}

	account, err := api.AccountService.GetAccount(r.Context(), int32(id))
	if err != nil {
		log.Printf("Error retrieving account: %v", err)
		utils.EncodeJSON(w, r, http.StatusInternalServerError, map[string]any{"error": "Failed to retrieve account"})
		return
	}
	if account.ID == 0 {
		log.Printf("Account not found for ID: %d", id)
		utils.EncodeJSON(w, r, http.StatusNotFound, map[string]any{"error": "Account not found"})
		return
	}
	utils.EncodeJSON(w, r, http.StatusOK, account)
}
