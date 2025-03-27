package utils

import "github.com/welibekov/grantmaster/internal/role/types"

// Diff computes the difference between two slices of Role structs.
// It returns a slice of Roles that are in the first slice but not in the second.
func Diff(roles, input []types.Role) []types.Role {
	return diffRoles(roles, input)
}

// diffRoles compares two slices of Role structs.
// For each Role in the first slice:
//   - If it is missing entirely in the second slice (by Name), include it fully in the result.
//   - If it exists in both, compare the Schemas and include only those differences.
func diffRoles(first, second []types.Role) []types.Role {
	// Create a lookup for the second slice of Roles by Name.
	secondMap := make(map[string]types.Role)
	for _, r := range second {
		secondMap[r.Name] = r
	}

	var diff []types.Role
	// Iterate over each Role in the first slice.
	for _, r1 := range first {
		if r2, ok := secondMap[r1.Name]; !ok {
			// Role missing in second; include entire Role.
			diff = append(diff, r1)
		} else {
			// Role exists in both; compare Schemas.
			schemasDiff := diffSchemas(r1.Schemas, r2.Schemas)
			if len(schemasDiff) > 0 {
				// Include Role with the differing Schemas.
				diff = append(diff, types.Role{
					Name:    r1.Name,
					Schemas: schemasDiff,
				})
			}
		}
	}
	return diff
}

// diffSchemas compares two slices of Schema structs.
// For each Schema in the first slice:
//   - If the Schema does not exist in the second slice (matched by Schema field), it is returned entirely.
//   - If the Schema exists in both, then its Grants are compared and only the differences (grants in first not in second) are returned.
func diffSchemas(first, second []types.Schema) []types.Schema {
	// Create a lookup for the second schemas by the Schema field.
	secondMap := make(map[string]types.Schema)
	for _, s := range second {
		secondMap[s.Schema] = s
	}

	var diff []types.Schema
	// Iterate over each Schema in the first slice.
	for _, s1 := range first {
		if s2, ok := secondMap[s1.Schema]; !ok {
			// The whole Schema from first is missing in second; include it.
			diff = append(diff, s1)
		} else {
			// Compare grants between the two Schemas.
			grantsDiff := diffGrants(s1.Grants, s2.Grants)
			if len(grantsDiff) > 0 {
				// Include Schema with differing Grants.
				diff = append(diff, types.Schema{
					Schema: s1.Schema,
					Grants: grantsDiff,
				})
			}
		}
	}
	return diff
}

// diffGrants compares two slices of grant strings.
// It returns a slice of grants that are present in the first slice but not in the second.
func diffGrants(first, second []string) []string {
	// Create a set for the second slice of grants for quick lookup.
	secondSet := make(map[string]bool)
	for _, grant := range second {
		secondSet[grant] = true
	}
	
	var diff []string
	// Iterate over each grant in the first slice.
	for _, grant := range first {
		if !secondSet[grant] {
			// Grant is present in first but missing in second; include it.
			diff = append(diff, grant)
		}
	}
	return diff
}
