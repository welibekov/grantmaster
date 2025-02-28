package policy

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/welibekov/grantmaster/modules/policy"
	"github.com/welibekov/grantmaster/modules/policy/types"
)

// Apply applies the provided policies to the system by revoking existing policies
// that are no longer applicable and granting new policies as necessary.
func (p *PGPolicy) Apply(ctx context.Context, policies []types.Policy) error {
	// Add a prefix related to roles to the incoming policies for consistency
	policies = p.addRolePrefix(policies)

	// Retrieve existing policies from the database to compare against
	exisitingPolicies, err := p.GetExisting(ctx)
	if err != nil {
		// Return an error if retrieving existing policies fails
		return fmt.Errorf("couldn't apply policies: %v", err)
	}

	// Determine which policies need to be revoked by comparing existing and new policies
	revokePolicies := policy.WhatToRevoke(policies, exisitingPolicies)

	// Log the count of policies identified for revocation for debugging purposes
	logrus.Debugln("Revoke policies length=", len(revokePolicies))

	// Revoke the identified policies from the database
	if err := p.Revoke(ctx, revokePolicies); err != nil {
		// Return an error if revocation of policies fails
		return fmt.Errorf("couldn't revoke policies: %v", err)
	}

	// Determine which new policies need to be granted by comparing existing and new policies
	grantPolicies := policy.WhatToGrant(policies, exisitingPolicies)

	// Grant the new policies to the database and return any potential error
	return p.Grant(ctx, grantPolicies)
}
