package policy

import "github.com/jackc/pgx/v5/pgxpool"

func New(pool *pgxpool.Pool) *PGPolicy {
	return &PGPolicy{
		pool: pool,
	}
}

type PGPolicy struct {
	pool *pgxpool.Pool
}
