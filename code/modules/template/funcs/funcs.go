package funcs

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"github.com/welibekov/grantmaster/modules/role/types"
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

func (f *Funcs) IsRoleExist(role types.Role) bool {
	query := fmt.Sprintf(`SELECT 1 FROM pg_roles WHERE rolname = '%s';`, role.Name)

	logrus.Debugln(query)

	rows, err := f.pool.Query(f.ctx, query)
	if err != nil {
		logrus.Errorf("couldn't find if role exist: %v", err)
		return false
	}
	defer rows.Close()

	var exist int

	for rows.Next() {
		if err := rows.Scan(&exist); err != nil {
			logrus.Errorf("failed to scan row for checking role: %v", err)
			return false
		}

		break
	}

	if err := rows.Err(); err != nil {
		logrus.Errorf("error occurred while iterating over rows: %v", err)
		return false
	}

	return exist == 1
}
