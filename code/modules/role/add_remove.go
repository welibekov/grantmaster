package role

import (
	"github.com/welibekov/grantmaster/modules/role/types"
)

// WhatToCreate determines which roles from slice A do not exist in slice B
// and returns a slice of those unique roles.
func WhatToCreate(A, B []types.Role) []types.Role {
	// Create a map to track roles in slice B for fast lookup
	roleMap := make(map[string]struct{})

	// Populate the map with roles from slice B
	for _, role := range B {
		roleMap[role.Name] = struct{}{}
	}

	var uniqueRoles []types.Role

	// Iterate through slice A and find roles that are not in slice B
	for _, role := range A {
		if _, exists := roleMap[role.Name]; !exists {
			uniqueRoles = append(uniqueRoles, role)
		}
	}

	return uniqueRoles
}

// WhatToRemove determines which roles from slice B do not exist in slice A
// and can be thought of as the opposite of WhatToCreate.
func WhatToRemove(A, B []types.Role) []types.Role {
	// Call WhatToCreate with reversed arguments to identify roles to remove
	return WhatToCreate(B, A)
}
