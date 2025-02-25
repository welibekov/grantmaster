package role

import "github.com/welibekov/grantmaster/modules/role/types"

func WhatToCreate(A, B []types.Role) []types.Role {
	roleMap := make(map[string]struct{})

	for _, role := range B {
		roleMap[role.Name] = struct{}{}
	}

	var uniqueRoles []types.Role

	for _, role := range A {
		if _, exists := roleMap[role.Name]; !exists {
			uniqueRoles = append(uniqueRoles, role)
		}
	}

	return uniqueRoles
}

func WhatToRemove(A, B []types.Role) []types.Role {
	return WhatToCreate(B, A)
}
