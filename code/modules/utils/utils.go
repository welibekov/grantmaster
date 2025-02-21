package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"

	"github.com/welibekov/grantmaster/modules/types"
)

func ReadPoliciesFromDirectory(policyDirectory string) ([]types.Policy, error) {
	var policies []types.Policy

	err := filepath.Walk(policyDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error accessing the path %q: %w", path, err)
		}

		// Check if the file has a .yaml extension
		if !info.IsDir() && (filepath.Ext(path) == ".yaml" || filepath.Ext(path) == ".yml") {
			// Read the file
			file, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("could not read file %q: %v", path, err)
			}

			// Unmarshal the YAML into a Policy struct
			var policy types.Policy
			err = yaml.Unmarshal(file, &policy)
			if err != nil {
				return fmt.Errorf("could not unmarshal YAML from file %q: %v", path, err)
			}

			// Add the policy to the slice
			policies = append(policies, policy)
		}

		return nil
	})

	if err != nil {
		return policies, fmt.Errorf("error walking the path %q: %w", policyDirectory, err)
	}

	return policies, nil
}
