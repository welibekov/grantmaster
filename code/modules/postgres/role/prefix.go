package role

import "github.com/welibekov/grantmaster/modules/role/types"

// addRolePrefix adds a predefined prefix to the name of each role in the provided slice.
// It returns the modified slice of roles with the updated names.
func (p *PGRole) addRolePrefix(roles []types.Role) []types.Role {
	// Iterate over each role in the slice using an index to allow modification.
	for index, role := range roles {
		// Prepend the role prefix to the role's name.
		role.Name = p.rolePrefix + role.Name

		// Update the role in the original slice with the modified role.
		roles[index] = role
	}

	// Return the modified slice of roles.
	return roles
}
