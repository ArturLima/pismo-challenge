package services

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/ArturLima/pismo/internal/store/pgstore"
	"github.com/ArturLima/pismo/test/mocks"
	"github.com/golang/mock/gomock"
)

func TestAccountService_CreateAccount_Success(t *testing.T) {
	result := pgstore.Account{ID: 1, Document: "1234"}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockQ := mocks.NewMockaccountQueries(ctrl)
	mockQ.
		EXPECT().
		CreateAccount(gomock.Any(), "1234").
		Return(result, nil)

	svc := NewAccountService(nil, mockQ)

	final, err := svc.CreateAccount(context.Background(), "1234")
	if err != nil {
		t.Fatalf("esperava sem erro, mas recebeu: %v", err)
	}

	if !reflect.DeepEqual(final, result) {
		t.Fatalf("esperava %v, mas recebeu %v", result, final)
	}
}

func TestAccountService_createAccount_error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockQ := mocks.NewMockaccountQueries(ctrl)

	mockQ.EXPECT().CreateAccount(gomock.Any(), "1234").Return(pgstore.Account{}, fmt.Errorf("db is down"))

	service := NewAccountService(nil, mockQ)

	_, err := service.CreateAccount(context.Background(), "1234")

	if err == nil {
		t.Fatal("expected an error, got nil")
	}

}

func TestAccountService_getAccount_success(t *testing.T) {
	result := pgstore.Account{ID: 1, Document: "1234"}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocksQ := mocks.NewMockaccountQueries(ctrl)

	mocksQ.EXPECT().GetAccountById(gomock.Any(), int32(1)).Return(result, nil)

	service := NewAccountService(nil, mocksQ)

	account, err := service.GetAccount(context.Background(), int32(1))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !reflect.DeepEqual(account, result) {
		t.Fatalf("expected account %v, got %v", result, account)
	}
}

func TestAccountService_getAccount_error(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockQ := mocks.NewMockaccountQueries(ctrl)

	mockQ.EXPECT().GetAccountById(gomock.Any(), int32(0)).Return(pgstore.Account{}, fmt.Errorf("Not found"))

	service := NewAccountService(nil, mockQ)

	_, err := service.GetAccount(context.Background(), 0)
	if err == nil {
		t.Fatal("expected an error, got nil")
	}

}
