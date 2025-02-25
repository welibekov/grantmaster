package role

import (
	"context"
	"fmt"

	"github.com/welibekov/grantmaster/modules/role/types"
)

// GetExisting retrieves existing roles from the database.
func (p *PGRole) GetExisting(ctx context.Context) ([]types.Role, error) {
	// Map to hold roles associated with each username.
	rolesMap := make(map[string][]types.Schema)

	// Slice to hold the resulting roles.
	roles := make([]types.Role, 0)

	query := fmt.Sprintf(`
SELECT
    nspname AS schema_name,
    rolname AS role_name,
    pg_catalog.has_schema_privilege(rolname, n.oid, 'USAGE') AS has_usage,
    pg_catalog.has_schema_privilege(rolname, n.oid, 'CREATE') AS has_create
FROM
    pg_catalog.pg_roles r
JOIN
    pg_catalog.pg_namespace n
    ON pg_catalog.has_schema_privilege(r.rolname, n.oid, 'USAGE')
WHERE
    r.rolname LIKE '%s%%'
    AND n.nspname NOT LIKE 'pg_%%'
    AND n.nspname != 'information_schema'
    AND n.nspname != 'pg_catalog'
    AND n.nspname != 'public'
ORDER BY
    schema_name, role_name;
`, p.rolePrefix)

	rows, err := p.pool.Query(ctx, query)
	if err != nil {
		// Wrap the error to include context about where it occurred.
		return roles, fmt.Errorf("failed to execute query to get policies: %w", err)
	}
	defer rows.Close() // Ensure rows are closed after processing.

	// Process each row returned by the query.
	for rows.Next() {
		var (
			schema, role        string
			hasUsage, hasCreate bool
		)

		// Scan the row into the username and role variables.
		if err := rows.Scan(&schema, &role, &hasUsage, &hasCreate); err != nil {
			// Wrap the error to include context about scanning the row.
			return roles, fmt.Errorf("failed to scan row for username: %w", err)
		}

		grants := p.getGrants(hasUsage, hasCreate)

		// FIXME: This is a only workaround and should be removed in future.
		// Please conduct meeting with database and datateam teams to
		// find proper solution.
		if len(grants) == 1 && grants[0] == "usage" { // workaround
			grants = append(grants, "select")
		}

		// Map the role to the corresponding username.
		rolesMap[role] = append(rolesMap[role], types.Schema{
			Schema: schema, Grants: append([]string{}, grants...)})
	}

	// Check for errors that may have occurred during the row iteration.
	if err := rows.Err(); err != nil {
		// Wrap the error to indicate an issue occurred while processing rows.
		return roles, fmt.Errorf("error occurred while iterating over rows: %w", err)
	}

	// Convert the rolesMap to a list of policies.
	for role, schemas := range rolesMap {
		roles = append(roles, types.Role{
			Name:    role,
			Schemas: schemas,
		})
	}

	// Return the list of policies and no error.
	return roles, nil
}

// getGrants generates a list of grant permissions based on the provided flags.
// It takes two boolean parameters indicating whether to include "usage" and/or "create" permissions.
//
// Parameters:
// - hasUsage: a boolean indicating if the "usage" permission should be included.
// - hasCreate: a boolean indicating if the "create" permission should be included.
//
// Returns:
// A slice of strings containing the granted permissions. It can include
// "usage", "create", or both, depending on the input parameters.
func (p *PGRole) getGrants(hasUsage, hasCreate bool) []string {
	// Initialize a slice to hold the grants with a capacity of 2.
	grants := make([]string, 0, 2)

	// Check if the "usage" permission should be added.
	if hasUsage {
		grants = append(grants, "usage")
	}

	// Check if the "create" permission should be added.
	if hasCreate {
		grants = append(grants, "create")
	}

	// Return the slice of grants.
	return grants
}
