package policy

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/welibekov/grantmaster/modules/policy/types"
)

func GetExisting(ctx context.Context, pool *pgxpool.Pool) ([]types.Policy, error) {
	rolesMap := make(map[string][]string)
	policies := make([]types.Policy, 0)

	query := `
SELECT u.usename
AS username, r.rolname
AS role
FROM pg_user u
JOIN pg_auth_members m
ON u.usesysid = m.member
JOIN pg_roles r
ON m.roleid = r.oid;
`
	rows, err := pool.Query(ctx, query)
	if err != nil {
		return policies, fmt.Errorf("couldn't run query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var username, role string

		if err := rows.Scan(&username, &role); err != nil {
			return policies, fmt.Errorf("couldn't scan row: %v", err)
		}

		rolesMap[username] = append(rolesMap[username], role)
	}

	if rows.Err() != nil {
		return policies, fmt.Errorf("couldn't not run query on rows: %v", err)
	}

	for username, roles := range rolesMap {
		policies = append(policies, types.Policy{
			Username: username,
			Roles:    roles,
		})
	}

	return policies, nil
}
