package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	_ "github.com/ArturLima/pismo/docs" // This is required to load the Swagger docs
	"github.com/ArturLima/pismo/internal/api"
	"github.com/ArturLima/pismo/internal/services"
	"github.com/ArturLima/pismo/internal/store/pgstore"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

// @title Pismo API
// @version 1.0
// @description pismo endpoint 2.
func main() {

	_ = godotenv.Load()
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		os.Getenv("PISMO_DATABASE_USER"),
		os.Getenv("PISMO_DATABASE_PASSWORD"),
		os.Getenv("PISMO_DATABASE_HOST"),
		os.Getenv("PISMO_DATABASE_PORT"),
		os.Getenv("PISMO_DATABASE_NAME"),
	))
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		panic(err)
	}

	accountSvc := services.NewAccountService(pool, pgstore.New(pool))
	transactionSvc := services.NewTransactionService(pool, pgstore.New(pool))

	api := api.Api{
		Router:             chi.NewMux(),
		AccountService:     accountSvc,
		TransactionService: transactionSvc,
	}

	api.BindRoutes()

	fmt.Println("Start server on port: 3080")
	if err := http.ListenAndServe(":3080", api.Router); err != nil {
		panic(err)
	}
}
