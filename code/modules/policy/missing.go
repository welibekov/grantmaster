package policy

import (
	"github.com/welibekov/grantmaster/modules/policy/types"
)

// WhatToRevoke compares two slices of Policies and returns the missing roles in A based on B.
func WhatToRevoke(A, B []types.Policy) []types.Policy {
	roleMapA := make(map[string]map[string]struct{})

	// Create a map for policies in A with roles
	for _, policy := range A {
		if _, exists := roleMapA[policy.Username]; !exists {
			roleMapA[policy.Username] = make(map[string]struct{})
		}
		for _, role := range policy.Roles {
			roleMapA[policy.Username][role] = struct{}{}
		}
	}

	var result []types.Policy

	// Compare and find missing roles
	for _, policyB := range B {
		// We only care about users present in B
		if rolesA, exists := roleMapA[policyB.Username]; exists {
			missingRoles := []string{}
			for _, role := range policyB.Roles {
				if _, inRolesA := rolesA[role]; !inRolesA {
					missingRoles = append(missingRoles, role)
				}
			}
			if len(missingRoles) > 0 {
				result = append(result, types.Policy{Username: policyB.Username, Roles: missingRoles})
			}
		} else {
			// If the username from B is not in A, include all roles from B
			result = append(result, policyB)
		}
	}

	return result
}
