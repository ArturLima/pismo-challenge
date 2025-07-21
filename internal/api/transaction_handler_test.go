package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/ArturLima/pismo/internal/store/pgstore"
	"github.com/ArturLima/pismo/internal/useCases/transaction"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
)

type TransactioNServiceFake struct {
	createTransactionFk func(ctx context.Context, req transaction.CreateTransactionRequest) (pgstore.Transaction, error)
}

func (t *TransactioNServiceFake) CreateTransaction(ctx context.Context, req transaction.CreateTransactionRequest) (pgstore.Transaction, error) {
	return t.createTransactionFk(ctx, req)
}

func TestCreateTransactionHandlerSuccess(t *testing.T) {
	var amt pgtype.Numeric
	err := amt.Scan("100.00")
	assert.NoError(t, err)

	api := &Api{
		TransactionService: &TransactioNServiceFake{
			createTransactionFk: func(ctx context.Context, req transaction.CreateTransactionRequest) (pgstore.Transaction, error) {
				return pgstore.Transaction{
					ID:          1,
					AccountID:   1,
					Amount:      amt,
					OperationID: 4,
				}, nil
			},
		},
	}

	payload := transaction.CreateTransactionRequest{
		AccountId:     1,
		OperationType: 4,
		Amount:        "100.00",
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest("POST", "/transactions", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	api.CreateTransaction(w, req)

	assert.Equal(t, 201, w.Code)

}
