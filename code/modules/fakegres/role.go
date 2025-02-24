package fakegres

import (
	"context"
	"fmt"
	"os"

	"github.com/welibekov/grantmaster/modules/role/types"
)

// ApplyRole processes a slice of roles and applies the specified actions (grant/revoke).
// It updates the roles map and ensures that any absent roles are removed.
func (f *Fakegres) ApplyRole(_ context.Context, roles []types.Role) error {
	updateRolesMap := make(map[string][]string) // Map to keep track of role updates

	// Loop through each role to apply it.
	for _, role := range roles {
		updateRolesMap[role.Name] = []string{} // Update the map

		// Apply the individual role.
		if err := apply(role, f.absPath(f.roleDir, role.Name)); err != nil {
			// Wrap and return the error with context about which role failed.
			return fmt.Errorf("failed to apply role for user %s: %w", role.Name, err)
		}
	}

	// Remove any roles that are not present in the updated roles map.
	if err := f.removeAbsentRoles(updateRolesMap); err != nil {
		// Wrap and return the error for better clarity on failure.
		return fmt.Errorf("failed to remove absent roles: %w", err)
	}

	return nil // Successfully applied all roles
}

// removeAbsentRoles removes roles for users that are not present in the updatePolicesMap.
// It compares the existing roles with the incoming updates and removes any roles
// for users that are no longer required.
func (f *Fakegres) removeAbsentRoles(updateRolesMap map[string][]string) error {
	// Read the existing roles from the storage.
	existingRolesMap, err := readExisting[types.Role](f.roleDir,
		func(roles []types.Role) map[string][]string {
			rolesMap := make(map[string][]string)
			for _, role := range roles {
				rolesMap[role.Name] = []string{}
			}

			return rolesMap
		},
	)
	if err != nil {
		// Wrap the error with additional context before returning.
		return fmt.Errorf("failed to read existing roles: %w", err)
	}

	// Iterate over existing roles to identify which ones need to be removed.
	for name := range existingRolesMap {
		_, found := updateRolesMap[name]
		if !found {
			// If the username is not found in the update map, proceed to remove the role.
			if err := os.Remove(f.absPath(f.roleDir, name)); err != nil {
				// Wrap the error with additional context before returning.
				return fmt.Errorf("failed to remove role '%s': %w", name, err)
			}
		}
	}

	return nil
}
