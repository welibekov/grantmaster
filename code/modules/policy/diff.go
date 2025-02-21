package policy

import (
	"github.com/welibekov/grantmaster/modules/policy/types"
)

// Function to compare two slices of Policy
func Compare(A []types.Policy, B []types.Policy) []types.Policy {
	rolesMap := make(map[string][]string)

	// Build a map from slice B
	for _, policyB := range B {
		rolesMap[policyB.Username] = policyB.Roles
	}

	var result []types.Policy

	// Compare policies in slice A with those in map from slice B
	for _, policyA := range A {
		rolesB, exists := rolesMap[policyA.Username]
		if !exists || !equalRoles(policyA.Roles, rolesB) {
			result = append(result, policyA)
		}
	}

	return result
}

// Helper function to compare two role slices
func equalRoles(rolesA, rolesB []string) bool {
	if len(rolesA) != len(rolesB) {
		return false
	}

	roleMap := make(map[string]bool)
	for _, role := range rolesB {
		roleMap[role] = true
	}

	for _, role := range rolesA {
		if !roleMap[role] {
			return false
		}
	}

	return true
}
