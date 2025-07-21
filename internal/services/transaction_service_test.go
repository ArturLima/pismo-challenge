package services

import (
	"context"
	"testing"

	"github.com/ArturLima/pismo/internal/store/pgstore"
	"github.com/ArturLima/pismo/internal/useCases/transaction"
	"github.com/jackc/pgx/v5/pgtype"
)

type mockTransactionService struct {
	createTransactionFake func(ctx context.Context, request pgstore.CreateTransactionParams) (pgstore.Transaction, error)
}

func (m *mockTransactionService) CreateTransaction(ctx context.Context, request pgstore.CreateTransactionParams) (pgstore.Transaction, error) {
	return m.createTransactionFake(ctx, request)
}

func TestTransactionService_CreateTransaction_OperationCreditVoucher_success(t *testing.T) {
	var amt pgtype.Numeric
	if err := amt.Scan("100"); err != nil {
		t.Fatalf("failed to scan amount: %v", err)
	}

	result := pgstore.Transaction{ID: 1, AccountID: 1, OperationID: 4, Amount: amt}

	service := NewTransactionService(nil, &mockTransactionService{
		createTransactionFake: func(ctx context.Context, request pgstore.CreateTransactionParams) (pgstore.Transaction, error) {
			if request.AccountID != 1 {
				t.Fatalf("expected AccountID 1, got %d", request.AccountID)
			}
			if request.OperationID != 4 {
				t.Fatalf("expected OperationID 4, got %d", request.OperationID)
			}
			resultFloat, _ := amt.Float64Value()
			resultWaited, _ := result.Amount.Float64Value()

			if resultFloat != resultWaited {
				t.Fatalf("expected Amount %v, got %v", amt, request.Amount)
			}
			return result, nil
		},
	})

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
