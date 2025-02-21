package fakegres

import (
	"fmt"

	"github.com/welibekov/grantmaster/modules/utils"
)

// readExistingPolicies reads all YAML files in the specified directory and unmarshals them into Policy structs.
func (f *Fakegres) readExistingPolicies() (map[string][]string, error) {
	policiesMap := make(map[string][]string)

	policies, err := utils.ReadPoliciesFromDirectory(f.rootDir)
	if err != nil {
		return nil, fmt.Errorf("error walking the path %q: %w", f.rootDir, err)
	}

	for _, policy := range policies {
		policiesMap[policy.Username] = policy.Roles
	}

	return policiesMap, nil
}
