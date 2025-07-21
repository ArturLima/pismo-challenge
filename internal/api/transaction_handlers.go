package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ArturLima/pismo/internal/useCases/transaction"
	"github.com/ArturLima/pismo/internal/utils"
)

// CreateTransaction creates a financial transaction for a given account
// @Summary Create Transaction
// @Description Create a new transaction (purchase, withdrawal, credit) for an account.
// @Tags transactions
// @Accept json
// @Produce json
// @Param transaction body transaction.CreateTransactionRequest true "Transaction data"
// @Success 201 {object} transaction.TransactionResponse
// @Failure 400 {object} map[string]any
// @Failure 500 {object} map[string]any
// @Router /transactions [post]
func (api *Api) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	tsc, err := utils.DecodeJSON[transaction.CreateTransactionRequest](r)
	if err != nil {
		log.Println("Failed to decode request body", err)
		utils.EncodeJSON(w, r, http.StatusBadRequest, map[string]any{"error": "Invalid request body"})
		return
	}

	amt, _ := strconv.ParseFloat(tsc.Amount, 64)

	if err := transaction.ValidateTransaction(tsc.OperationType, amt); err != nil {
		log.Println("Validation error", err)
		utils.EncodeJSON(w, r, http.StatusBadRequest, map[string]any{"error": err.Error()})
		return
	}

	result, err := api.TransactionService.CreateTransaction(r.Context(), tsc)
	if err != nil {
		log.Println("Failed to create transaction", err)
		utils.EncodeJSON(w, r, http.StatusInternalServerError, map[string]any{"error": "Failed to create transaction"})
		return
	}

	utils.EncodeJSON(w, r, http.StatusCreated, transaction.ToTransacationResponse(result))

	fmt.Println(tsc)

}
