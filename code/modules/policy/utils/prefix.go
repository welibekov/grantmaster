package utils

import (
	"fmt"
	"strings"

	"github.com/welibekov/grantmaster/modules/policy/types"
)

// CheckPrefix checks whether all roles in the provided policies start
// with the given prefix. If any role does not start with the prefix,
// an error is returned indicating which role failed the check.
//
// Parameters:
// - policies: a slice of Policy structs that contain the roles to check.
// - prefix: the prefix that each role must start with.
//
// Returns:
// - An error if any role does not start with the prefix, otherwise returns nil.
func CheckPrefix(policies []types.Policy, prefix string) error {
	for _, policy := range policies {
		for _, role := range policy.Roles {
			if !strings.HasPrefix(role, prefix) {
				return fmt.Errorf("role %s doesn't start with prefix %s", role, prefix)
			}
		}
	}

	return nil
}
