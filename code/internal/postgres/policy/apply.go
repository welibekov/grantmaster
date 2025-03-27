package policy

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/welibekov/grantmaster/internal/policy/types"
	"github.com/welibekov/grantmaster/internal/policy/utils"
	"github.com/welibekov/grantmaster/internal/utils/debug"
)

// Apply applies the provided policies to the system by revoking existing policies
// that are no longer applicable and granting new policies as necessary.
func (p *PGPolicy) Apply(ctx context.Context, policies []types.Policy) error {
	defer p.pool.Close() // Close the connection pool when the function completes

	// Retrieve existing policies from the database to compare against
	existingPolicies, err := p.Get(ctx)
	if err != nil {
		// Return an error if retrieving existing policies fails
		return fmt.Errorf("couldn't apply policies: %v", err)
	}

	// Output the existing policies for debugging purposes
	debug.OutputMarshal(existingPolicies, "existing policies")

	// Determine which policies need to be revoked by comparing existing and new policies
	revokePolicies := utils.Diff(policies, existingPolicies)
	// Output the revoke policies for debugging purposes
	debug.OutputMarshal(revokePolicies, "revoke policies")

	// Determine which new policies need to be granted by comparing existing and new policies
	grantPolicies := utils.Diff(existingPolicies, policies)
	// Output the grant policies for debugging purposes
	debug.OutputMarshal(grantPolicies, "grant policies")

	// Log the count of policies identified for revocation for debugging purposes
	logrus.Debugln("Revoke policies length=", len(revokePolicies))

	// Revoke the identified policies from the database
	if err := p.Revoke(ctx, revokePolicies); err != nil {
		// Return an error if revocation of policies fails
		return fmt.Errorf("couldn't revoke policies: %v", err)
	}

	// Grant the new policies to the database and return any potential error
	return p.Grant(ctx, grantPolicies)
}
