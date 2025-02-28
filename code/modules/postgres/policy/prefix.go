package policy

import "github.com/welibekov/grantmaster/modules/policy/types"

// PGPolicy struct is assumed to be defined elsewhere in your package
type PGPolicy struct {
	rolePrefix string // rolePrefix is the prefix to be added to each role in the policies
}

// addRolePrefix adds a prefix to each role in the given slice of policies.
// It iterates through each policy and appends the rolePrefix to every role.
func (p *PGPolicy) addRolePrefix(policies []types.Policy) []types.Policy {
	// Iterate over each policy in the policies slice
	for policyIndex, policy := range policies {
		// Iterate over each role in the current policy
		for roleIndex, role := range policy.Roles {
			// Add the predefined prefix to the role
			policy.Roles[roleIndex] = p.rolePrefix + role
		}

		// Update the modified policy back into the policies slice
		policies[policyIndex] = policy
	}

	// Return the modified slice of policies
	return policies
}
