package role

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/welibekov/grantmaster/modules/role/types"
)

func (p *PGRole) Grant(ctx context.Context, roles []types.Role) error {
	var query string

	for _, role := range roles {
		exist, err := p.IsRoleExist(ctx, role)
		if err != nil {
			return fmt.Errorf("couldn't check if role %s exist: %v", role.Name, err)
		}

		if !exist {
			query += fmt.Sprintf("CREATE ROLE %s;", role.Name)
		}

		for _, schema := range role.Schemas {
			tablesExistInSchema := p.TablesExistInSchema(ctx, schema)

			for _, grant := range schema.Grants {
				if p.IsTableLevelGrant(grant) {

					if !tablesExistInSchema {
						logrus.Warnf("no tables in schema '%s', skipping '%s' grant", schema.Schema, grant)
						continue
					}

					query += fmt.Sprintf("GRANT %s ON ALL TABLES IN SCHEMA %s TO %s;", grant, schema.Schema, role.Name)
				} else {
					query += fmt.Sprintf("GRANT %s ON SCHEMA %s TO %s;", grant, schema.Schema, role.Name)
				}
			}
		}
	}

	logrus.Debugln(query) // Log the generated query for debugging purposes

	_, err := p.pool.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("couldn't execute grant query '%s': %v", query, err)
	}

	return nil
}
