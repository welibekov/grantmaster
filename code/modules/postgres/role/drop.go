package role

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/welibekov/grantmaster/modules/role/types"
)

// Drop removes the specified roles and revokes their privileges from the schemas.
func (p *PGRole) Drop(ctx context.Context, roles []types.Role) error {
	var query string // Initialize an empty string to build the SQL query.

	// Iterate over each role to process drop and revoke operations.
	for _, role := range roles {
		// Check if the current role exists in the database.
		exist, err := p.IsRoleExist(ctx, role)
		if err != nil {
			// Return an error if the existence check fails.
			return fmt.Errorf("couldn't check if role %s exist: %v", role.Name, err)
		}

		// If the role does not exist, skip further operations for this role.
		if !exist { 
			continue
		}

		// Revoke privileges from all schemas associated with the role.
		for _, schema := range role.Schemas {
			query += fmt.Sprintf("REVOKE ALL PRIVILEGES ON SCHEMA %s FROM %s;", schema.Schema, role.Name)
			query += fmt.Sprintf("REVOKE GRANT OPTION FOR ALL PRIVILEGES ON SCHEMA %s FROM %s;", schema.Schema, role.Name)
			query += fmt.Sprintf("REVOKE ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA %s FROM %s;", schema.Schema, role.Name)
			query += fmt.Sprintf("REVOKE ALL PRIVILEGES ON ALL FUNCTIONS IN SCHEMA %s FROM %s;", schema.Schema, role.Name)
		}

		// Construct the drop role query for the current role.
		query += fmt.Sprintf("DROP ROLE %s;", role.Name)

		// Log the generated query for debugging purposes.
		logrus.Debugln(query)
	}

	// Execute the constructed SQL query to drop roles and revoke privileges.
	_, err := p.pool.Exec(ctx, query)
	if err != nil {
		// Return an error if the execution of the drop query fails.
		return fmt.Errorf("couldn't execute drop query '%s': %v", query, err)
	}

	// Return nil, indicating success.
	return nil
}
