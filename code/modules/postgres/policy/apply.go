package policy

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/welibekov/grantmaster/modules/policy"
	"github.com/welibekov/grantmaster/modules/policy/types"
)

func (p *PGPolicy) Apply(ctx context.Context, policies []types.Policy) error {
	// Retrieve existing policies from the database
	exisitingPolicies, err := p.GetExisting(ctx)
	if err != nil {
		return fmt.Errorf("couldn't apply policies: %v", err)
	}

	// Determine which policies need to be revoked based on the current and new policies
	revokePolicies := policy.WhatToRevoke(policies, exisitingPolicies)

	// Log the length of policies to be revoked for debugging purposes
	logrus.Debugln("Revoke policies length=", len(revokePolicies))

	// Revoke the identified policies from the database
	if err := p.Revoke(ctx, revokePolicies); err != nil {
		return fmt.Errorf("couldn't revoke policies: %v", err)
	}

	// Determine which policies need to be granted based on the current and new policies
	grantPolicies := policy.WhatToGrant(policies, exisitingPolicies)

	// Grant the new policies to the database
	return p.Grant(ctx, grantPolicies)
}
