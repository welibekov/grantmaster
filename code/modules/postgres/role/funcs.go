package role

import (
	"context"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/welibekov/grantmaster/modules/role/types"
	"github.com/welibekov/grantmaster/modules/utils"
)

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
