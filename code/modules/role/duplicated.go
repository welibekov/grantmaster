package role

import (
	"fmt"

	"github.com/welibekov/grantmaster/modules/role/types"
)

// DetectDuplicated checks for duplicate roles based on the user.
// It returns an error if a duplicate is found.
func DetectDuplicated(roles []types.Role) error {
	// Create a map to track users that have already been processed.
	rolesMap := map[string]struct{}{}

	// Iterate over each role in the provided slice.
	for _, role := range roles {
		// Check if the user of the current role already exists in the map.
		_, found := rolesMap[role.Name]
		if found {
			// If the user is found, return an error indicating a duplicate.
			return fmt.Errorf("duplicated role for '%s' user", role.Name)
		}

		// If the user is not found, add it to the map.
		rolesMap[role.Name] = struct{}{}
	}

	// If no duplicates were found, return nil indicating success.
	return nil
}
