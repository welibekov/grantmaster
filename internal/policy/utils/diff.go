package utils

import "github.com/welibekov/grantmaster/internal/policy/types"

// DiffPolicies returns a slice of Policy representing what exists in candPolicies
// that is missing or different in refPolicies. For each policy in candPolicies:
//   - If the policy (by Username) is missing in refPolicies, the entire policy is included.
//   - If it exists, only the extra roles (in candPolicies, not in refPolicies) are included.
func Diff(refPolicies, candPolicies []types.Policy) []types.Policy {
	// Build a lookup map for reference policies by Username.
	refMap := make(map[string]types.Policy)
	for _, p := range refPolicies {
		refMap[p.Username] = p // Store each reference policy by its Username.
	}

	var diff []types.Policy // Slice to hold the policies that differ or are missing.

	for _, candPolicy := range candPolicies {
		if refPolicy, ok := refMap[candPolicy.Username]; !ok {
			// This policy is new in the candidate because it does not exist in reference.
			diff = append(diff, candPolicy)
		} else {
			// Compare roles: get roles that are in candidate but not in reference.
			rolesDiff := diffRoles(refPolicy.Roles, candPolicy.Roles)
			if len(rolesDiff) > 0 {
				// Create a new policy with the differing roles and add it to the diff slice.
				diff = append(diff, types.Policy{
					Username: candPolicy.Username,
					Roles:    rolesDiff,
				})
			}
		}
	}
	return diff // Return the slice of policies that are different or missing.
}

// diffRoles returns the roles present in candRoles that do not appear in refRoles.
func diffRoles(refRoles, candRoles []string) []string {
	refSet := make(map[string]bool) // Create a set for quick lookup of reference roles.
	for _, role := range refRoles {
		refSet[role] = true // Add each role from refRoles to the set.
	}

	var diff []string // Slice to hold roles that are different.

	for _, role := range candRoles {
		if !refSet[role] {
			// If the role is not found in the reference set, add it to the diff slice.
			diff = append(diff, role)
		}
	}
	return diff // Return the list of roles that differ.
}
