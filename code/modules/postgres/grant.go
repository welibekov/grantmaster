package postgres

import (
	"context"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/welibekov/grantmaster/modules/policy/types"
)

// grantPolicy takes a context and a slice of policies, and attempts to grant specified roles to users
func (p *Postgres) grantPolicy(ctx context.Context, policies []types.Policy) error {
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

// grantQuery constructs the SQL query string for granting roles to a user
func (p *Postgres) grantQuery(policy types.Policy) string {
	// Join the roles with a comma and format the SQL grant statement
	return fmt.Sprintf(`GRANT %s TO %s;`, strings.Join(policy.Roles, ","), policy.Username)
}
