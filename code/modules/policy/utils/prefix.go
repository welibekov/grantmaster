package utils

import (
	"fmt"
	"strings"

	"github.com/welibekov/grantmaster/modules/policy/types"
)

func CheckPrefix(policies []types.Policy, prefix string) error {
	for _, policy := range policies {
		for _, role := range policy.Roles {
			if !strings.HasPrefix(role, prefix) {
				return fmt.Errorf("role %s doesn't starts with prefix %s", role, prefix)
			}
		}
	}

	return nil
}
