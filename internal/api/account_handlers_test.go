package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ArturLima/pismo/internal/store/pgstore"
	"github.com/ArturLima/pismo/internal/useCases/account"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

type fakeAccountService struct {
	createAccountfake func(ctx context.Context, document string) (pgstore.Account, error)
	getAccountFake    func(ctx context.Context, id int32) (pgstore.Account, error)
}

func (f *fakeAccountService) CreateAccount(ctx context.Context, document string) (pgstore.Account, error) {
	return f.createAccountfake(ctx, document)
}

func (f *fakeAccountService) GetAccount(ctx context.Context, id int32) (pgstore.Account, error) {
	return f.getAccountFake(ctx, id)
}

func TestCreateAccountHandler_Success(t *testing.T) {
	api := &Api{
		AccountService: &fakeAccountService{
			createAccountfake: func(ctx context.Context, document string) (pgstore.Account, error) {
				return pgstore.Account{
					ID: 1, Document: document}, nil
			},
		},
	}

	payload := map[string]string{"document": "12345678901"}
	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/accounts", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	api.CreateAccount(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var resp account.AccountResponse
	err := json.NewDecoder(rr.Body).Decode(&resp)
	assert.NoError(t, err)

}

func TestCreateAccountHandler_WithBodyWrong_Erro(t *testing.T) {

	api := &Api{
		AccountService: &fakeAccountService{},
	}

	req := httptest.NewRequest(http.MethodPost, "/accounts", bytes.NewReader([]byte("{")))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	api.CreateAccount(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)

}

func TestCreateAccountHandler_WithInternaLServerError(t *testing.T) {

	api := &Api{
		AccountService: &fakeAccountService{
			createAccountfake: func(ctx context.Context, document string) (pgstore.Account, error) {
				return pgstore.Account{}, errors.New("internal server error")
			},
		},
	}

	payload := map[string]string{"document": "12345678901"}
	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/accounts", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	api.CreateAccount(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)

}

func TestGetAccountHandler_success(t *testing.T) {

	result := pgstore.Account{ID: 1, Document: "12345678901"}

	api := &Api{
		AccountService: &fakeAccountService{
			getAccountFake: func(ctx context.Context, id int32) (pgstore.Account, error) {
				assert.Equal(t, int32(1), id)
				return pgstore.Account{ID: 1, Document: "12345678901"}, nil
			},
		},
	}

	req := httptest.NewRequest(http.MethodGet, "/accounts/1", nil)
	rr := httptest.NewRecorder()

	r := chi.NewRouter()

	r.Get("/accounts/{accountId}", api.GetAccount)
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var resp account.AccountResponse
	err := json.NewDecoder(rr.Body).Decode(&resp)
	assert.NoError(t, err)
	assert.Equal(t, result.ID, resp.ID)

}

func TestGetAccountHandler_BadRequest(t *testing.T) {

	api := &Api{
		AccountService: &fakeAccountService{
			getAccountFake: func(ctx context.Context, id int32) (pgstore.Account, error) {
				return pgstore.Account{ID: 1, Document: "12345678901"}, nil
			},
		},
	}

	req := httptest.NewRequest(http.MethodGet, "/accounts/abc", nil)
	rr := httptest.NewRecorder()
	r := chi.NewRouter()
	r.Get("/accounts/{accountId}", api.GetAccount)
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}
