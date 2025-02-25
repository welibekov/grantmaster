package role

import (
	"context"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/welibekov/grantmaster/modules/role/types"
)

func (p *PGRole) Create(ctx context.Context, roles []types.Role) error {
	for _, role := range roles {
		query := p.createQuery(role)

		logrus.Debugln(query) // Log the generated query for debugging purposes

		// Execute the revoke query on the PostgreSQL database
		_, err := p.pool.Exec(ctx, query)
		if err != nil {
			// Wrap the error with more context to help identify the issue if it occurs
			return fmt.Errorf("couldn't revoke grants from role'%s': %w", role.Name, err)
		}
	}

	// Return nil if all policies were processed without errors
	return nil
}

func (p *PGRole) createQuery(role types.Role) string {
	var query string

	for _, schema := range role.Schemas {
		query += fmt.Sprintf(`CREATE ROLE %s;`, role.Name)
		for _, grant := range schema.Grants {
			if strings.ToUpper(grant) == "SELECT" {
				query += fmt.Sprintf(`GRANT %s ON ALL TABLES IN SCHEMA %s TO %s`, grant, schema.Schema, role.Name)
			} else {
				query += fmt.Sprintf(`GRANT %s ON %s TO %s;`, grant, schema.Schema, role.Name)
			}
		}
	}

	return query
}
