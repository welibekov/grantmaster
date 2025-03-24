package policy

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/welibekov/grantmaster/modules/policy/types"
)

// Get retrieves existing policies from the database based on the specified RolePrefix.
func (p *PGPolicy) Get(ctx context.Context) ([]types.Policy, error) {
	// Map to hold roles associated with each username.
	rolesMap := make(map[string][]string)

	// Slice to hold the resulting policies.
	policies := make([]types.Policy, 0)

	// SQL query to fetch usernames and their associated roles, filtered by RolePrefix.
	query := fmt.Sprintf(`
SELECT u.usename AS username, r.rolname AS role
FROM pg_user u
JOIN pg_auth_members m ON u.usesysid = m.member
JOIN pg_roles r ON m.roleid = r.oid
WHERE r.rolname LIKE '%s%%';
`, p.RolePrefix)

	// Log the SQL query for debugging purposes.
	logrus.Debugln("policy get query:", query)

	// Execute the SQL query against the database.
	rows, err := p.pool.Query(ctx, query)
	if err != nil {
		// Return an error if the query execution fails, wrapping the original error for context.
		return policies, fmt.Errorf("failed to execute query to get policies: %w", err)
	}
	defer rows.Close() // Ensure the rows are closed after processing to prevent memory leaks.

	// Process each row returned by the query.
	for rows.Next() {
		var username, role string

		// Scan the row into the username and role variables.
		if err := rows.Scan(&username, &role); err != nil {
			// Return an error if scanning the row fails, wrapping the original error for context.
			return policies, fmt.Errorf("failed to scan row for username: %w", err)
		}

		// Map the role to the corresponding username.
		rolesMap[username] = append(rolesMap[username], role)
	}

	// Check for errors that may have occurred during the row iteration.
	if err := rows.Err(); err != nil {
		// Return an error indicating an issue that occurred while processing rows.
		return policies, fmt.Errorf("error occurred while iterating over rows: %w", err)
	}

	// Convert the rolesMap to a list of policies.
	for username, roles := range rolesMap {
		policies = append(policies, types.Policy{
			Username: username, // Set the username for the policy.
			Roles:    roles,    // Set the associated roles for the policy.
		})
	}

	// Return the list of policies and no error.
	return policies, nil
}
