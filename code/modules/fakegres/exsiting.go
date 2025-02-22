package fakegres

import (
	"fmt"

	"github.com/welibekov/grantmaster/modules/policy"
	"github.com/welibekov/grantmaster/modules/policy/types"
)

// readExistingPolicies reads all YAML files in the specified directory and unmarshals them into Policy structs.
func (f *Fakegres) readExistingPolicies() (map[string][]string, error) {
	policiesMap := make(map[string][]string)

	policies, err := policy.ReadPoliciesFromDirectory(f.rootDir,
		func(path string) ([]types.Policy, error) {
			policies := []types.Policy{}

			policy, err := policy.ReadPolicy[types.Policy](path)
			if err != nil {
				return policies, err
			}

			return append(policies, policy), nil
		})

	if err != nil {
		return nil, fmt.Errorf("error walking the path %q: %w", f.rootDir, err)
	}

	for _, policy := range policies {
		policiesMap[policy.Username] = policy.Roles
	}

	return policiesMap, nil
}
