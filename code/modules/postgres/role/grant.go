package role

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/welibekov/grantmaster/modules/role/types"
)

// Grant issues grants for the specified roles. It creates roles if they do not exist
// and assigns the corresponding schema grants.
func (p *PGRole) Grant(ctx context.Context, roles []types.Role) error {
	var query string // Initialize an empty query string to hold the SQL commands

	for _, role := range roles {
		// Check if the role already exists in the database
		exist, err := p.IsRoleExist(ctx, role)
		if err != nil {
			return fmt.Errorf("couldn't check if role %s exists: %v", role.Name, err)
		}

		// If the role does not exist, create it
		if !exist {
			query += fmt.Sprintf("CREATE ROLE %s;", role.Name)
		}

		// Iterate over each schema associated with the role
		for _, schema := range role.Schemas {
			// Check if there are tables in the schema
			tablesExistInSchema := p.IsTablesExistInSchema(ctx, schema)

			// Iterate over the grants specified for the schema
			for _, grant := range schema.Grants {
				if p.IsItTableLevelGrant(grant) {
					// If it's a table-level grant and there are no tables in the schema, log a warning
					if !tablesExistInSchema {
						logrus.Warnf("no tables in schema '%s', skipping '%s' grant", schema.Schema, grant)
						continue // Skip to the next grant if there are no tables
					}

					// Construct the grant command for all tables in the schema
					query += fmt.Sprintf("GRANT %s ON ALL TABLES IN SCHEMA %s TO %s;", grant, schema.Schema, role.Name)
				} else {
					// Construct the grant command for the schema itself
					query += fmt.Sprintf("GRANT %s ON SCHEMA %s TO %s;", grant, schema.Schema, role.Name)
				}
			}
		}
	}

	logrus.Debugln(query) // Log the generated query for debugging purposes

	// Execute the assembled query to apply the role and grants in the database
	_, err := p.pool.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("couldn't execute grant query '%s': %v", query, err)
	}

	return nil // Return nil if no errors occurred
}
