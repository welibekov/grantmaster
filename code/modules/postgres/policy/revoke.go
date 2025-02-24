package policy

import (
	"context"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/welibekov/grantmaster/modules/policy/types"
)

// Revoke revokes specified roles from a user based on the provided policies.
// It takes a context and a slice of policies as input and returns an error if any issues occur during execution.
func (p *PGPolicy) Revoke(ctx context.Context, policies []types.Policy) error {
	// Iterate through each policy to revoke roles from the corresponding user
	for _, policy := range policies {
		// Generate the SQL revoke query for the current policy
		query := p.revokeQuery(policy)

		logrus.Debugln(query) // Log the generated query for debugging purposes

		// Execute the revoke query on the PostgreSQL database
		_, err := p.pool.Exec(ctx, query)
		if err != nil {
			// Wrap the error with more context to help identify the issue if it occurs
			return fmt.Errorf("couldn't revoke roles from user '%s': %w", policy.Username, err)
		}
	}

	// Return nil if all policies were processed without errors
	return nil
}

// revokeQuery constructs a SQL REVOKE statement for the given policy,
// which specifies roles to be revoked from a specific user.
func (p *PGPolicy) revokeQuery(policy types.Policy) string {
	// Wrap roles with double quotes
	roles := make([]string, 0, len(policy.Roles))

	for _, role := range policy.Roles {
		roles = append(roles, fmt.Sprintf(`"%s"`, role))
	}

	// Join the roles with commas and format the SQL query string
	return fmt.Sprintf(`REVOKE %s FROM "%s"`, strings.Join(roles, ","), policy.Username)
}
