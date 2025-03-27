package utils

import (
	"reflect"
	"sort"

	"github.com/welibekov/grantmaster/internal/role/types"
)

// Equal checks if two slices of roles are equal.
// It normalizes both slices and compares them using reflect.DeepEqual.
func Equal(roleA, roleB []types.Role) bool {
	roleA = normalizeRoles(roleA) // Normalize the first role slice
	roleB = normalizeRoles(roleB) // Normalize the second role slice

	// Use DeepEqual to determine if both normalized role slices are identical.
	return reflect.DeepEqual(roleA, roleB)
}

// normalizeRoles takes a slice of roles and normalizes them for comparison.
// It sorts the Schemas of each role by Schema name and
// sorts the Grants within each Schema.
func normalizeRoles(roles []types.Role) []types.Role {
	rolesSorted := []types.Role{} // Initialize a new slice to hold the sorted roles

	for _, role := range roles {
		// Sort Schemas by Schema name for consistency in comparison.
		sort.Slice(role.Schemas, func(i, j int) bool {
			return role.Schemas[i].Schema < role.Schemas[j].Schema
		})

		// Sort Grants within each Schema to ensure the order does not affect equality.
		for i := range role.Schemas {
			sort.Strings(role.Schemas[i].Grants) // Sort the Grants slice within each Schema
		}

		// Append the normalized role to the list of sorted roles.
		rolesSorted = append(rolesSorted, role)
	}

	// Sort the entire slice of roles by role name for consistent ordering.
	sort.Slice(rolesSorted, func(i, j int) bool {
		return rolesSorted[i].Name < rolesSorted[j].Name
	})

	// Return the normalized list of roles.
	return rolesSorted
}
