package main

import (
	storage "AllReady/internal/storage"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

const psqlDSN = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

func main() {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, psqlDSN)
	if err != nil {
		log.Fatal(err)
	}

	pgRepository := storage.NewPgRepository(pool)
	defer pool.Close()

	pgRepository.FillAllTables(ctx)
}
