package role

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/welibekov/grantmaster/internal/role/types"
)

// Revoke revokes specified roles from their associated schemas and grants.
// It checks if each role exists before attempting to revoke access.
// The revocation is performed for all tables and schemas defined in the roles.
func (p *PGRole) Revoke(ctx context.Context, roles []types.Role) error {
	var query string

	// Iterate over each role to process revocation.
	for _, role := range roles {
		// Check if the current role exists in the database.
		exist, err := p.IsRoleExist(ctx, role)
		if err != nil {
			return fmt.Errorf("couldn't check if role %s exists: %v", role.Name, err)
		}

		// Skip revocation for non-existing roles.
		if !exist {
			continue
		}

		// Iterate over the schemas associated with the role.
		for _, schema := range role.Schemas {
			// Iterate over the grants associated with the schema.
			for _, grant := range schema.Grants {
				// Construct the REVOKE query based on the type of grant (table-level or schema-level).
				if p.IsItTableLevelGrant(grant) {
					query += fmt.Sprintf("REVOKE %s ON ALL TABLES IN SCHEMA %s FROM %s;", grant, schema.Schema, role.Name)
				} else {
					query += fmt.Sprintf("REVOKE %s ON SCHEMA %s FROM %s;", grant, schema.Schema, role.Name)
				}
			}
		}
	}

	// Log the generated query for debugging purposes.
	logrus.Debugln(query) 

	// Execute the compiled REVOKE query.
	_, err := p.pool.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("couldn't execute drop query '%s': %v", query, err)
	}

	return nil // Return nil if the revocation was successful.
}
