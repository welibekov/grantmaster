package role

import "github.com/welibekov/grantmaster/modules/role/types"

func Diff(roles, input []types.Role) []types.Role {
	return diffRoles(roles, input)
}

// diffRoles compares two slices of Role structs.
// For each Role in the first slice:
//   - If it is missing entirely in the second slice (by Name), include it fully.
//   - If it exists in both, compare the Schemas and include only those differences.
func diffRoles(first, second []types.Role) []types.Role {
	// Create a lookup for second Roles by Name.
	secondMap := make(map[string]types.Role)
	for _, r := range second {
		secondMap[r.Name] = r
	}

	var diff []types.Role
	for _, r1 := range first {
		if r2, ok := secondMap[r1.Name]; !ok {
			// Role missing in second; include entire Role.
			diff = append(diff, r1)
		} else {
			// Role exists in both; compare Schemas.
			schemasDiff := diffSchemas(r1.Schemas, r2.Schemas)
			if len(schemasDiff) > 0 {
				diff = append(diff, types.Role{
					Name:    r1.Name,
					Schemas: schemasDiff,
				})
			}
		}
	}
	return diff
}

// diffSchemas compares two []Schema slices.
// For each Schema in first:
//   - If the schema does not exist in second (matched by Schema field), it is returned entirely.
//   - If the schema exists in both, then its Grants are compared and only the differences (grants in first not in second) are returned.
func diffSchemas(first, second []types.Schema) []types.Schema {
	// Create a lookup for the second schemas by the Schema field.
	secondMap := make(map[string]types.Schema)
	for _, s := range second {
		secondMap[s.Schema] = s
	}

	var diff []types.Schema
	for _, s1 := range first {
		if s2, ok := secondMap[s1.Schema]; !ok {
			// The whole schema from first is missing in second.
			diff = append(diff, s1)
		} else {
			// Compare grants.
			grantsDiff := diffGrants(s1.Grants, s2.Grants)
			if len(grantsDiff) > 0 {
				diff = append(diff, types.Schema{
					Schema: s1.Schema,
					Grants: grantsDiff,
				})
			}
		}
	}
	return diff
}

// diffGrants returns a slice of grants that are present in first but not in second.
func diffGrants(first, second []string) []string {
	secondSet := make(map[string]bool)
	for _, grant := range second {
		secondSet[grant] = true
	}
	var diff []string
	for _, grant := range first {
		if !secondSet[grant] {
			diff = append(diff, grant)
		}
	}
	return diff
}
