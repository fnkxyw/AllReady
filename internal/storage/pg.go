package storage

import "github.com/jackc/pgx/v5/pgxpool"

type PgRepository struct {
	pool *pgxpool.Pool
}

func NewPgRepository(p *pgxpool.Pool) PgRepository {
	return PgRepository{pool: p}
}
