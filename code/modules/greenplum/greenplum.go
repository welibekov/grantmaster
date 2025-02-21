package greenplum

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/welibekov/grantmaster/modules/database/base"
	"github.com/welibekov/grantmaster/modules/types"
)

type Greenplum struct {
	*base.Database

	pool       *pgxpool.Pool // connection pool
	connString string        // connection string
}

func New(config map[string]string) (*Greenplum, error) {
	connString, found := config["GM_GREENPLUM_CONN_STRING"]
	if !found {
		return nil, fmt.Errorf("GM_GREENPLUM_CONN_STRING not defined")
	}

	return &Greenplum{
		Database:   base.NewDatabase(),
		connString: connString,
	}, nil
}

func (g *Greenplum) ApplyPolicy(ctx context.Context, policies []types.Policy) error {
	pool, err := pgxpool.New(ctx, g.connString)
	if err != nil {
		return fmt.Errorf("couldn't connect to database: %v", err)
	}
	defer pool.Close()

	g.pool = pool

	return fmt.Errorf("NYI")
}
