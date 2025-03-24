package utils

import (
	"fmt"
	"strings"

	"github.com/welibekov/grantmaster/modules/role/types"
)

// CheckPrefix checks if all roles in the provided slice have names that start with the specified prefix.
// It returns an error if any role does not match the prefix.
func CheckPrefix(roles []types.Role, prefix string) error {
	for _, role := range roles {
		// Check if the role name starts with the given prefix
		if !strings.HasPrefix(role.Name, prefix) {
			// Return an error if it doesn't
			return fmt.Errorf("role %s doesn't start with prefix %s", role.Name, prefix)
		}
	}

	// Return nil if all roles match the prefix
	return nil
}
