package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ArturLima/pismo/internal/store/pgstore"
	"github.com/ArturLima/pismo/internal/useCases/transaction"
	"github.com/ArturLima/pismo/test/mocks"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
)

func TestCreateTransactionHandlerSuccess(t *testing.T) {

	var amt pgtype.Numeric
	err := amt.Scan("100.00")
	assert.NoError(t, err)

	payload := transaction.CreateTransactionRequest{
		AccountId:     1,
		OperationType: 4,
		Amount:        "100.00",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockITransactionService(ctrl)

	service.EXPECT().CreateTransaction(gomock.Any(), payload).Return(pgstore.Transaction{ID: 1, OperationID: 4, AccountID: 1, Amount: amt}, nil)

	handler := &Api{
		TransactionService: service,
	}

	body, _ := json.Marshal(payload)
	req := httptest.NewRequest("POST", "/transactions", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.CreateTransaction(rr, req)

	if rr.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", rr.Code)
	}
}
