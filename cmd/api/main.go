package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/ArturLima/pismo/internal/api"
	"github.com/ArturLima/pismo/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
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

	api := api.Api{
		Router:             chi.NewMux(),
		AccountService:     services.NewAccountService(pool),
		TransactionService: services.NewTransactionService(pool),
	}

	api.BindRoutes()

	fmt.Println("Start server on port: 3080")
	if err := http.ListenAndServe(":3080", api.Router); err != nil {
		panic(err)
	}
}
