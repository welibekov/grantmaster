package role

import (
	"context"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/welibekov/grantmaster/modules/role/types"
	"github.com/welibekov/grantmaster/modules/utils"
)

func (p *PGRole) IsExist(ctx context.Context, query string) (bool, error) {
	logrus.Debugln(query)

	rows, err := p.pool.Query(ctx, query)
	if err != nil {
		return false, fmt.Errorf("couldn't run query: %v", err)
	}
	defer rows.Close()

	var exist int

	for rows.Next() {
		if err := rows.Scan(&exist); err != nil {
			return false, fmt.Errorf("failed to scan row for 'exist' value: %v", err)
		}

		break
	}

	if err := rows.Err(); err != nil {
		return false, fmt.Errorf("error occurred while iterating over rows: %v", err)
	}

	return exist == 1, nil
}

func (p *PGRole) IsRoleExist(ctx context.Context, role types.Role) (bool, error) {
	query := fmt.Sprintf(`SELECT 1 FROM pg_roles WHERE rolname = '%s';`, role.Name)

	logrus.Debugln(query)

	rows, err := p.pool.Query(ctx, query)
	if err != nil {
		return false, fmt.Errorf("couldn't find if role exist: %v", err)
	}
	defer rows.Close()

	var exist int

	for rows.Next() {
		if err := rows.Scan(&exist); err != nil {
			return false, fmt.Errorf("failed to scan row for checking role: %v", err)
		}

		break
	}

	if err := rows.Err(); err != nil {
		return false, fmt.Errorf("error occurred while iterating over rows: %v", err)
	}

	return exist == 1, nil
}

func (p *PGRole) IsTableLevelGrant(grant string) bool {
	return utils.In(strings.ToUpper(grant), TableLevelGrants)
}

func (p *PGRole) TablesExistInSchema(ctx context.Context, schema types.Schema) bool {
	query := fmt.Sprintf(`
SELECT COUNT(*) 
FROM information_schema.tables
WHERE table_schema = '%s' AND table_type = 'BASE TABLE';
`, schema.Schema)

	exist, err := p.IsExist(ctx, query)
	if err != nil {
		logrus.Warnf("checking tables existance in schema %s failed: %v", schema.Schema, err)
		return false
	}

	return exist
}
