package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ArturLima/pismo/internal/useCases/transaction"
	"github.com/ArturLima/pismo/internal/utils"
)

func (api *Api) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	transaction, err := utils.DecodeJSON[transaction.CreateTransactionRequest](r)
	if err != nil {
		log.Println("Failed to decode request body", err)
		utils.EncodeJSON(w, r, http.StatusBadRequest, map[string]any{"error": "Invalid request body"})
		return
	}

	result, err := api.TransactionService.CreateTransaction(r.Context(), transaction)
	if err != nil {
		log.Println("Failed to create transaction", err)
		utils.EncodeJSON(w, r, http.StatusInternalServerError, map[string]any{"error": "Failed to create transaction"})
		return
	}

	utils.EncodeJSON(w, r, http.StatusCreated, result)

	fmt.Println(transaction)

}
