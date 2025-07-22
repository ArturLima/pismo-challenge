package services

import (
	"context"
	"sync"

	"github.com/ArturLima/pismo/internal/store/pgstore"
	"github.com/ArturLima/pismo/internal/useCases/transaction"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type transactionQueries interface {
	CreateTransaction(ctx context.Context, request pgstore.CreateTransactionParams) (pgstore.Transaction, error)
}

type ITransactionService interface {
	CreateTransaction(ctx context.Context, request transaction.CreateTransactionRequest) (pgstore.Transaction, error)
}

type TransactionService struct {
	mu      sync.Mutex
	pool    *pgxpool.Pool
	queries transactionQueries
}

func NewTransactionService(pool *pgxpool.Pool, q transactionQueries) ITransactionService {
	return &TransactionService{
		pool:    pool,
		queries: q,
	}
}

func (s *TransactionService) CreateTransaction(ctx context.Context, request transaction.CreateTransactionRequest) (pgstore.Transaction, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var pgAmount pgtype.Numeric
	if err := pgAmount.Scan(request.Amount); err != nil {
		return pgstore.Transaction{}, err
	}

	transaction := pgstore.CreateTransactionParams{
		AccountID:   int32(request.AccountId),
		OperationID: int32(request.OperationType),
		Amount:      pgAmount,
	}

	ts, err := s.queries.CreateTransaction(ctx, transaction)
	if err != nil {
		return pgstore.Transaction{}, err
	}
	return ts, nil

}
