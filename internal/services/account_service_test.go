package services

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/ArturLima/pismo/internal/store/pgstore"
)

type mockAccountQueries struct {
	createAccountFake  func(ctx context.Context, document string) (pgstore.Account, error)
	getAccountByIdFake func(ctx context.Context, id int32) (pgstore.Account, error)
}

func (m *mockAccountQueries) CreateAccount(ctx context.Context, document string) (pgstore.Account, error) {
	return m.createAccountFake(ctx, document)
}

func (m *mockAccountQueries) GetAccountById(ctx context.Context, id int32) (pgstore.Account, error) {
	return m.getAccountByIdFake(ctx, id)
}

func TestAccountService_createAccoutn_success(t *testing.T) {
	result := pgstore.Account{ID: 1, Document: "1234"}

	service := NewAccountService(nil, &mockAccountQueries{
		createAccountFake: func(ctx context.Context, document string) (pgstore.Account, error) {
			if document != "1234" {
				t.Fatalf("expected document '1234', got '%s'", document)
			}
			return result, nil
		},
	})

	account, err := service.CreateAccount(context.Background(), "1234")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !reflect.DeepEqual(account, result) {
		t.Fatalf("expected account %v, got %v", result, account)
	}
}

func TestAccountService_createAccount_error(t *testing.T) {

	service := NewAccountService(nil, &mockAccountQueries{
		createAccountFake: func(ctx context.Context, document string) (pgstore.Account, error) {
			return pgstore.Account{}, fmt.Errorf("error creating account")
		},
	})

	_, err := service.CreateAccount(context.Background(), "1234")
	if err == nil {
		t.Fatal("expected an error, got nil")
	}

}

func TestAccountService_getAccount_success(t *testing.T) {
	result := pgstore.Account{ID: 1, Document: "1234"}

	service := NewAccountService(nil, &mockAccountQueries{
		getAccountByIdFake: func(ctx context.Context, id int32) (pgstore.Account, error) {
			if id != 1 {
				t.Fatalf("expected id 1, got %d", id)
			}
			return result, nil
		},
	})

	account, err := service.GetAccount(context.Background(), 1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !reflect.DeepEqual(account, result) {
		t.Fatalf("expected account %v, got %v", result, account)
	}

}

func TestAccountService_getAccount_error(t *testing.T) {
	service := NewAccountService(nil, &mockAccountQueries{
		getAccountByIdFake: func(ctx context.Context, id int32) (pgstore.Account, error) {
			return pgstore.Account{}, fmt.Errorf("error getting account")
		},
	})

	_, err := service.GetAccount(context.Background(), 1)
	if err == nil {
		t.Fatal("expected an error, got nil")
	}

}
