package services

import (
	"context"

	"github.com/ArturLima/pismo/internal/store/pgstore"
	"github.com/ArturLima/pismo/internal/useCases/transaction"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ITransactionService interface {
	CreateTransaction(ctx context.Context, request transaction.CreateTransactionRequest) (pgstore.Transaction, error)
}

type TransactionService struct {
	pool    *pgxpool.Pool
	queries *pgstore.Queries
}

func NewTransactionService(pool *pgxpool.Pool) ITransactionService {
	return &TransactionService{
		pool:    pool,
		queries: pgstore.New(pool),
	}
}

func (s *TransactionService) CreateTransaction(ctx context.Context, request transaction.CreateTransactionRequest) (pgstore.Transaction, error) {
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
