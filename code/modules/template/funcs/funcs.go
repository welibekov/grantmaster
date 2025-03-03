package funcs

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Funcs struct {
	ctx  context.Context
	pool *pgxpool.Pool
}

func New(ctx context.Context, pool *pgxpool.Pool) *Funcs {
	return &Funcs{
		ctx:  ctx,
		pool: pool,
	}
}
