package policy

import (
	"context"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/welibekov/grantmaster/internal/policy/types"
)

// Grant takes a context and a slice of policies, and attempts to grant specified roles to users.
// It iterates over each policy, generating and executing SQL grant queries for the roles defined in each policy.
// If any error occurs during the execution, it returns a wrapped error providing context about the failed operation.
func (p *PGPolicy) Grant(ctx context.Context, policies []types.Policy) error {
	// Iterate over each policy in the provided slice
	for _, policy := range policies {
		// Generate the SQL grant query based on the current policy
		query := p.grantQuery(policy)

		// Log the generated SQL query for debugging purposes
		logrus.Debugln(query)

		// Check if the policy specifies any roles to grant
		if len(policy.Roles) > 0 {
			// Execute the SQL grant query
			_, err := p.pool.Exec(ctx, query)
			if err != nil {
				// Wrap the error with additional context and return it
				return fmt.Errorf("couldn't grant roles %v for user '%s': %w", policy.Roles, policy.Username, err)
			}
		}
	}

	// Return nil if all roles have been granted successfully
	return nil
}

// grantQuery constructs the SQL query string for granting roles to a user.
// It takes a policy as input, formats the GRANT statement and returns it as a string.
func (p *PGPolicy) grantQuery(policy types.Policy) string {
	// Join the roles with a comma and format the SQL grant statement
	return fmt.Sprintf(`GRANT %s TO %s;`, strings.Join(policy.Roles, ","), policy.Username)
}
