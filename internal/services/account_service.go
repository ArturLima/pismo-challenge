package services

import (
	"context"

	"github.com/ArturLima/pismo/internal/store/pgstore"
	"github.com/jackc/pgx/v5/pgxpool"
)

type IAccountService interface {
	CreateAccount(ctx context.Context, document string) (pgstore.Account, error)
	GetAccount(ctx context.Context, accountId int32) (pgstore.Account, error)
}

type AccountService struct {
	pool    *pgxpool.Pool
	queries *pgstore.Queries
}

func NewAccountService(pool *pgxpool.Pool) IAccountService {
	return &AccountService{
		pool:    pool,
		queries: pgstore.New(pool),
	}
}

func (a *AccountService) CreateAccount(ctx context.Context, document string) (pgstore.Account, error) {
	account, err := a.queries.CreateAccount(ctx, document)
	if err != nil {
		return pgstore.Account{}, err
	}
	return account, nil
}

func (a *AccountService) GetAccount(ctx context.Context, accountId int32) (pgstore.Account, error) {
	account, err := a.queries.GetAccountById(ctx, accountId)
	if err != nil {
		return pgstore.Account{}, err
	}
	return account, nil
}
