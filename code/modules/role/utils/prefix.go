package utils

import (
	"fmt"
	"strings"

	"github.com/welibekov/grantmaster/modules/role/types"
)

func CheckPrefix(roles []types.Role, prefix string) error {
	for _, role := range roles {
		if !strings.HasPrefix(role.Name, prefix) {
			return fmt.Errorf("role %s doesn't starts with prefix %s", role.Name, prefix)
		}
	}

	return nil
}
