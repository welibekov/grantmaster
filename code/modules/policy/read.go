package policy

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"

	"github.com/welibekov/grantmaster/modules/policy/types"
	"github.com/welibekov/grantmaster/modules/utils"
)

// ReadPolicies reads policies from the specified path. If the path is a file and is
// a YAML file, it reads the policy directly. Otherwise, it attempts to read policies
// from all YAML files within a directory.
func ReadPolicies(policyPath string) ([]types.Policy, error) {
	if utils.IsItFile(policyPath) && utils.IsItYAML(policyPath) {
		// Read a single policy from a YAML file
		return readPolicyGen[[]types.Policy](policyPath)
	}

	// Read policies from all YAML files in the specified directory
	return readPoliciesFromDirectory(policyPath,
		func(path string) ([]types.Policy, error) {
			return readPolicyGen[[]types.Policy](path)
		})
}

// ReadPoliciesFromDirectory reads policies from all files in a directory using
// the provided policyFunc to handle each file.
func ReadPoliciesFromDirectory(policyPath string, policyFunc func(string) ([]types.Policy, error)) ([]types.Policy, error) {
	return readPoliciesFromDirectory(policyPath, policyFunc)
}

// ReadPolicy reads a single policy from the given file path and returns it as type T.
func ReadPolicy[T any](path string) (T, error) {
	return readPolicyGen[T](path)
}

// readPoliciesFromDirectory walks through the specified directory, applying policyFunc to read
// each YAML file and collecting the resulting policies.
func readPoliciesFromDirectory(
	policyPath string,
	policyFunc func(string) ([]types.Policy, error),
) ([]types.Policy, error) {
	var policies []types.Policy

	err := filepath.Walk(policyPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			// Wrap the error with additional context
			return fmt.Errorf("error accessing the path %q: %w", path, err)
		}

		// Check if the file has a .yaml or .yml extension
		if !info.IsDir() && utils.IsItYAML(path) {
			policy, err := policyFunc(path)
			if err != nil {
				// Wrap the error with context indicating which file failed
				return fmt.Errorf("could not read policy from file %q: %v", path, err)
			}

			// Add the retrieved policies to the slice
			policies = append(policies, policy...)
		}

		return nil
	})

	if err != nil {
		// Wrap the error indicating there was a problem walking the directory
		return policies, fmt.Errorf("error walking the path %q: %w", policyPath, err)
	}

	return policies, nil
}

// readPolicyGen reads the file at the given path and unmarshals its YAML content into the provided type T.
func readPolicyGen[T any](path string) (T, error) {
	var policy T

	// Read the file contents
	file, err := os.ReadFile(path)
	if err != nil {
		// Wrap the error to provide context
		return policy, fmt.Errorf("could not read file %q: %v", path, err)
	}

	// Unmarshal the YAML file content into the specified type
	if err := yaml.Unmarshal(file, &policy); err != nil {
		// Wrap the unmarshalling error with context
		return policy, fmt.Errorf("could not unmarshal YAML from file %q: %v", path, err)
	}

	return policy, nil
}
