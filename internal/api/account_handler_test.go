package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ArturLima/pismo/internal/store/pgstore"
	"github.com/ArturLima/pismo/test/mocks"
	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
)

func Test_GetAccountById(t *testing.T) {

	req := httptest.NewRequest("GET", "/accounts/1", nil)
	rr := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockIAccountService(ctrl)

	service.EXPECT().GetAccount(gomock.Any(), int32(1)).Return(pgstore.Account{
		ID:       1,
		Document: "12345678900",
	}, nil)

	// 2) Injeta no Api
	handler := &Api{AccountService: service}

	router := chi.NewRouter()
	router.Get("/accounts/{accountId}", handler.GetAccount)
	router.ServeHTTP(rr, req)

	if rr.Code != 200 {
		t.Errorf("Expected status code 200, got %d", rr.Code)
	}
}

func Test_CreateAccount(t *testing.T) {
	payload := map[string]string{
		"document": "12345678900"}
	body, _ := json.Marshal(payload)
	req := httptest.NewRequest("POST", "/accounts", bytes.NewReader(body))
	rr := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockIAccountService(ctrl)
	service.EXPECT().CreateAccount(gomock.Any(), "12345678900").Return(pgstore.Account{
		ID:       1,
		Document: "12345678900",
	}, nil)

	handler := &Api{AccountService: service}

	router := chi.NewRouter()
	router.Post("/accounts", handler.CreateAccount)
	router.ServeHTTP(rr, req)

	if rr.Code != 201 {
		t.Errorf("Expected status code 201, got %d", rr.Code)
	}
}

func Test_CreateAccount_withWrongPayload(t *testing.T) {
	req := httptest.NewRequest("POST", "/accounts", bytes.NewReader([]byte("{")))
	rr := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockIAccountService(ctrl)

	service.EXPECT().CreateAccount(gomock.Any(), gomock.Any()).Times(0)

	handler := &Api{AccountService: service}

	router := chi.NewRouter()
	router.Post("/accounts", handler.CreateAccount)
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %d", rr.Code)
	}

}

func Test_CreateAccount_WithInternalServerError(t *testing.T) {

	payload := map[string]string{"document": "123"}
	body, _ := json.Marshal(payload)
	req := httptest.NewRequest("POST", "/accounts", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockIAccountService(ctrl)

	service.EXPECT().CreateAccount(gomock.Any(), "123").Return(pgstore.Account{}, fmt.Errorf("db down"))

	handler := &Api{AccountService: service}

	router := chi.NewRouter()
	router.Post("/accounts", handler.CreateAccount)
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code 500, got %d", rr.Code)
	}

}
