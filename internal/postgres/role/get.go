package role

import (
	"context"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/welibekov/grantmaster/internal/role/types"
)

// Get retrieves existing roles and their grants from the database.
func (p *PGRole) Get(ctx context.Context) ([]types.Role, error) {
	// rolesMap maps each role to its corresponding schemas and their grants.
	rolesMap := make(map[string]map[string][]string)

	// roles will hold the resulting role objects to be returned.
	roles := make([]types.Role, 0)

	// Construct the SQL query to get roles along with their permissions.
	query := p.getQuery()

	logrus.Debugln("role get query:", query)

	// Execute the constructed query against the database.
	rows, err := p.pool.Query(ctx, query)
	if err != nil {
		// Wrap the error to include context about the failure.
		return roles, fmt.Errorf("failed to execute query to get roles: %w", err)
	}
	defer rows.Close() // Ensure rows are closed after processing.

	// Iterate through the result set.
	for rows.Next() {
		var (
			role, schema, permission, hasAccess string
		)

		// Scan the row into the variables for further processing.
		if err := rows.Scan(&role, &schema, &permission, &hasAccess); err != nil {
			// Wrap the error to indicate an issue during scanning.
			return roles, fmt.Errorf("failed to scan row: %w", err)
		}

		// Skip this entry if the role does not have access.
		if hasAccess != "YES" {
			continue
		}

		// Initialize the map entry for the role if it doesn't exist.
		if _, exist := rolesMap[role]; !exist {
			rolesMap[role] = make(map[string][]string)
		}

		// Initialize the map entry for the schema if it doesn't exist.
		if _, found := rolesMap[role][schema]; !found {
			rolesMap[role][schema] = []string{}
		}

		// Append the permission to the list of grants for the schema.
		rolesMap[role][schema] = append(rolesMap[role][schema], strings.ToLower(permission))
	}

	// Check if any errors occurred during the iteration of rows.
	if err := rows.Err(); err != nil {
		// Wrap the error to indicate an issue occurred while processing rows.
		return roles, fmt.Errorf("error occurred while iterating over rows: %w", err)
	}

	// Convert the rolesMap to a list of role objects for output.
	for role, schemasMap := range rolesMap {
		var schemas []types.Schema

		// Iterate over each schema in the map for the role.
		for schema, grants := range schemasMap {
			// Only include schemas that have grants.
			if len(grants) > 0 {
				schemas = append(schemas, types.Schema{
					Schema: schema,
					Grants: grants,
				})
			}
		}

		// Create a Role object and append it to the roles slice.
		roles = append(roles, types.Role{
			Name:    role,
			Schemas: schemas,
		})
	}

	// Return the list of roles with their associated schemas and no error.
	return roles, nil
}
