package storage

import "github.com/jackc/pgx/v5/pgxpool"

//repository struct / pgxpool v5

type PgRepository struct {
	pool *pgxpool.Pool
}

//constructor

func NewPgRepository(p *pgxpool.Pool) PgRepository {
	return PgRepository{pool: p}
}
