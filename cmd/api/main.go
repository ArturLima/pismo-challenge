package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/ArturLima/pismo/internal/api"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

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
		Router: chi.NewMux(),
	}

	api.BindRoutes()

	fmt.Println("Start server on port: 3080")
	if err := http.ListenAndServe("localhost:3080", nil); err != nil {
		panic(err)
	}
}
