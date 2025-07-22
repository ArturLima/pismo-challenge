package services

import (
	"context"
	"sync"

	"github.com/ArturLima/pismo/internal/store/pgstore"
	"github.com/jackc/pgx/v5/pgxpool"
)

type accountQueries interface {
	CreateAccount(ctx context.Context, document string) (pgstore.Account, error)
	GetAccountById(ctx context.Context, id int32) (pgstore.Account, error)
}

type IAccountService interface {
	CreateAccount(ctx context.Context, document string) (pgstore.Account, error)
	GetAccount(ctx context.Context, accountId int32) (pgstore.Account, error)
}

type AccountService struct {
	mu      sync.Mutex
	pool    *pgxpool.Pool
	queries accountQueries
}

func NewAccountService(pool *pgxpool.Pool, q accountQueries) IAccountService {
	return &AccountService{
		pool:    pool,
		queries: q,
	}
}

func (a *AccountService) CreateAccount(ctx context.Context, document string) (pgstore.Account, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	account, err := a.queries.CreateAccount(ctx, document)
	if err != nil {
		return pgstore.Account{}, err
	}
	return account, nil
}

func (a *AccountService) GetAccount(ctx context.Context, accountId int32) (pgstore.Account, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	account, err := a.queries.GetAccountById(ctx, accountId)
	if err != nil {
		return pgstore.Account{}, err
	}
	return account, nil
}
