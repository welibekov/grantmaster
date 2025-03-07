package role

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/welibekov/grantmaster/modules/role/types"
)

func (p *PGRole) Revoke(ctx context.Context, roles []types.Role) error {
	var query string

	for _, role := range roles {
		exist, err := p.IsRoleExist(ctx, role)
		if err != nil {
			return fmt.Errorf("couldn't check if role %s exist: %v", role.Name, err)
		}

		if !exist { // don't try to revoke grants on non-existing schema.
			continue
		}

		for _, schema := range role.Schemas {
			for _, grant := range schema.Grants {
				if p.IsTableLevelGrant(grant) {
					query += fmt.Sprintf("REVOKE %s ON ALL TABLES IN SCHEMA %s FROM %s;", grant, schema.Schema, role.Name)
				} else {
					query += fmt.Sprintf("REVOKE %s ON SCHEMA %s FROM %s;", grant, schema.Schema, role.Name)
				}
			}
		}
	}

	logrus.Debugln(query) // Log the generated query for debugging purposes

	_, err := p.pool.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("couldn't execute drop query '%s': %v", query, err)
	}

	return nil
}
