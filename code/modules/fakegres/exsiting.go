package fakegres

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/welibekov/grantmaster/modules/types"
	"gopkg.in/yaml.v3"
)

// readExistingPolicies reads all YAML files in the specified directory and unmarshals them into Policy structs.
func (f *Fakegres) readExistingPolicies() (map[string][]string, error) {
	var (
		policies    []types.Policy
		policiesMap = make(map[string][]string)
	)

	// Walk the directory
	err := filepath.Walk(f.rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error accessing the path %q: %w", path, err)
		}

		// Check if the file has a .yaml extension
		if !info.IsDir() && filepath.Ext(path) == ".yaml" {
			// Read the file
			file, err := os.ReadFile(path)
			if err != nil {
				logrus.Warnf("could not read file %q: %v", path, err)
				return nil // Continue to the next file
			}

			// Unmarshal the YAML into a Policy struct
			var policy types.Policy
			err = yaml.Unmarshal(file, &policy)
			if err != nil {
				logrus.Warnf("could not unmarshal YAML from file %q: %v", path, err)
				return nil // Continue to the next file
			}

			// Add the policy to the slice
			policies = append(policies, policy)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error walking the path %q: %w", f.rootDir, err)
	}

	for _, policy := range policies {
		policiesMap[policy.Username] = policy.Roles
	}

	return policiesMap, nil
}
