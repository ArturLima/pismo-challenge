package services

import (
	"context"
	"testing"

	"github.com/ArturLima/pismo/internal/store/pgstore"
	"github.com/ArturLima/pismo/internal/useCases/transaction"
	"github.com/ArturLima/pismo/test/mocks"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v5/pgtype"
)

func TestTransactionService_CreateTransaction_OperationCreditVoucher_success(t *testing.T) {
	var amt pgtype.Numeric
	if err := amt.Scan("100"); err != nil {
		t.Fatalf("failed to scan amount: %v", err)
	}

	result := pgstore.Transaction{ID: 1, AccountID: 1, OperationID: 4, Amount: amt}
	ctrl := gomock.NewController(t)

	service := mocks.NewMockITransactionService(ctrl)

	service.EXPECT().CreateTransaction(gomock.Any(), transaction.CreateTransactionRequest{
		AccountId:     1,
		OperationType: 4,
		Amount:        "100",
	}).Return(result, nil)

	transaction, err := service.CreateTransaction(context.Background(), transaction.CreateTransactionRequest{
		AccountId:     1,
		OperationType: 4,
		Amount:        "100",
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if transaction.ID != result.ID || transaction.AccountID != result.AccountID || transaction.OperationID != result.OperationID || transaction.Amount != result.Amount {
		t.Fatalf("expected transaction %v, got %v", result, transaction)
	}

}
