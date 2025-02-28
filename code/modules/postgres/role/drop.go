package role

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/welibekov/grantmaster/modules/role/types"
)

func (p *PGRole) Drop(ctx context.Context, roles []types.Role) error {
	for _, role := range roles {
		query := p.dropQuery(role)

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

func (p *PGRole) dropQuery(role types.Role) string {
	var query string
	for _, schema := range role.Schemas {
		query += fmt.Sprintf(`REVOKE ALL PRIVILEGES ON SCHEMA %s FROM %s;`, schema.Schema, role.Name)
		query += fmt.Sprintf(`REVOKE GRANT OPTION FOR ALL PRIVILEGES ON SCHEMA %s FROM %s;`, schema.Schema, role.Name)
	}

	query += fmt.Sprintf(`DROP ROLE %s`, role.Name)

	return query
}
