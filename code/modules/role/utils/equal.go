package utils

import (
	"reflect"
	"sort"

	"github.com/welibekov/grantmaster/modules/role/types"
)

// Equal checks if two slices of roles are equal.
// It normalizes both slices and compares them using reflect.DeepEqual.
func Equal(roleA, roleB []types.Role) bool {
	roleA = normalizeRoles(roleA)
	roleB = normalizeRoles(roleB)

	// Use DeepEqual to determine if both normalized role slices are identical.
	return reflect.DeepEqual(roleA, roleB)
}

// normalizeRoles takes a slice of roles and normalizes them for comparison.
// It sorts the Schemas of each role by Schema name and
// sorts the Grants within each Schema.
func normalizeRoles(roles []types.Role) []types.Role {
	rolesSorted := []types.Role{}

	for _, role := range roles {
		// Sort Schemas by Schema name for consistency in comparison.
		sort.Slice(role.Schemas, func(i, j int) bool {
			return role.Schemas[i].Schema < role.Schemas[j].Schema
		})

		// Sort Grants within each Schema to ensure the order does not affect equality.
		for i := range role.Schemas {
			sort.Strings(role.Schemas[i].Grants)
		}

		// Append the normalized role to the list of roleMaps.
		rolesSorted = append(rolesSorted, role)
	}

	// Return the normalized list of roles.
	return rolesSorted
}
