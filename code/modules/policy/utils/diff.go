package utils

import "github.com/welibekov/grantmaster/modules/policy/types"

// DiffPolicies returns a slice of Policy representing what exists in candPolicies
// that is missing or different in refPolicies. For each policy in candPolicies:
//   - If the policy (by Username) is missing in refPolicies, the entire policy is included.
//   - If it exists, only the extra roles (in candPolicies, not in refPolicies) are included.
func Diff(refPolicies, candPolicies []types.Policy) []types.Policy {
	// Build a lookup map for reference policies by Username.
	refMap := make(map[string]types.Policy)
	for _, p := range refPolicies {
		refMap[p.Username] = p
	}
	var diff []types.Policy
	for _, candPolicy := range candPolicies {
		if refPolicy, ok := refMap[candPolicy.Username]; !ok {
			// This policy is new in the candidate.
			diff = append(diff, candPolicy)
		} else {
			// Compare roles: get roles that are in candidate but not in reference.
			rolesDiff := diffRoles(refPolicy.Roles, candPolicy.Roles)
			if len(rolesDiff) > 0 {
				diff = append(diff, types.Policy{
					Username: candPolicy.Username,
					Roles:    rolesDiff,
				})
			}
		}
	}
	return diff
}

// diffRoles returns the roles present in candRoles that do not appear in refRoles.
func diffRoles(refRoles, candRoles []string) []string {
	refSet := make(map[string]bool)
	for _, role := range refRoles {
		refSet[role] = true
	}
	var diff []string
	for _, role := range candRoles {
		if !refSet[role] {
			diff = append(diff, role)
		}
	}
	return diff
}
