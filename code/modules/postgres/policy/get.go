package policy

import (
	"context"
	"fmt"

	"github.com/welibekov/grantmaster/modules/policy/types"
)

// GetExisting retrieves existing policies from the database.
func (p *PGPolicy) Get(ctx context.Context) ([]types.Policy, error) {
	// Map to hold roles associated with each username.
	rolesMap := make(map[string][]string)

	// Slice to hold the resulting policies.
	policies := make([]types.Policy, 0)

	// SQL query to fetch usernames and their associated roles.
	query := fmt.Sprintf(`
SELECT u.usename AS username, r.rolname AS role
FROM pg_user u
JOIN pg_auth_members m ON u.usesysid = m.member
JOIN pg_roles r ON m.roleid = r.oid
WHERE r.rolname LIKE '%s%%';
`, p.RolePrefix)

	// Execute the query.
	rows, err := p.pool.Query(ctx, query)
	if err != nil {
		// Wrap the error to include context about where it occurred.
		return policies, fmt.Errorf("failed to execute query to get policies: %w", err)
	}
	defer rows.Close() // Ensure rows are closed after processing.

	// Process each row returned by the query.
	for rows.Next() {
		var username, role string

		// Scan the row into the username and role variables.
		if err := rows.Scan(&username, &role); err != nil {
			// Wrap the error to include context about scanning the row.
			return policies, fmt.Errorf("failed to scan row for username: %w", err)
		}

		// Map the role to the corresponding username.
		rolesMap[username] = append(rolesMap[username], role)
	}

	// Check for errors that may have occurred during the row iteration.
	if err := rows.Err(); err != nil {
		// Wrap the error to indicate an issue occurred while processing rows.
		return policies, fmt.Errorf("error occurred while iterating over rows: %w", err)
	}

	// Convert the rolesMap to a list of policies.
	for username, roles := range rolesMap {
		policies = append(policies, types.Policy{
			Username: username,
			Roles:    roles,
		})
	}

	// Return the list of policies and no error.
	return policies, nil
}
