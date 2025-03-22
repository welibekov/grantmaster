package utils

import (
	"reflect"
	"sort"

	"github.com/welibekov/grantmaster/modules/role/types"
)

// Equal checks if two slices of roles are equal.
// It normalizes both slices and compares them using reflect.DeepEqual.
func Equal(roleA, roleB []types.Role) bool {
	roleAMap := normalizeRoles(roleA)
	roleBMap := normalizeRoles(roleB)

	// Use DeepEqual to determine if both normalized role slices are identical.
	return reflect.DeepEqual(roleAMap, roleBMap)
}

// normalizeRoles takes a slice of roles and normalizes them for comparison.
// It sorts the Schemas of each role by Schema name and
// sorts the Grants within each Schema.
func normalizeRoles(roles []types.Role) []types.Role {
	roleMaps := []types.Role{}

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
		roleMaps = append(roleMaps, role)
	}

	// Return the normalized list of roles.
	return roleMaps
}
