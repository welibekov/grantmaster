package utils

import (
	"github.com/welibekov/grantmaster/internal/role/types"
)

// WhatToGrant determines which roles from slice A do not exist in slice B
// and returns a slice of those unique roles.
// It leverages a map for efficient lookups to determine if a role from A is missing in B.
func WhatToGrant(A, B []types.Role) []types.Role {
	// Create a map to track roles in slice B for fast lookup
	roleMap := make(map[string]struct{})

	// Populate the map with roles from slice B
	for _, role := range B {
		roleMap[role.Name] = struct{}{}
	}

	// Initialize a slice to hold unique roles found in A but not in B
	var uniqueRoles []types.Role

	// Iterate through slice A and find roles that are not in slice B
	for _, role := range A {
		if _, exists := roleMap[role.Name]; !exists {
			// If the role is not found in the map, append it to uniqueRoles
			uniqueRoles = append(uniqueRoles, role)
		}
	}

	// Return the slice of roles unique to A
	return uniqueRoles
}

// WhatToDrop determines which roles from slice B do not exist in slice A
// and can be thought of as the opposite of WhatToGrant.
// It effectively identifies roles that are present in B but absent in A.
func WhatToDrop(A, B []types.Role) []types.Role {
	// Call WhatToGrant with reversed arguments to identify roles to remove
	return WhatToGrant(B, A)
}
