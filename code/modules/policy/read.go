package policy

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"

	"github.com/welibekov/grantmaster/modules/policy/types"
	"github.com/welibekov/grantmaster/modules/utils"
)

// ReadPolicies reads policies from the given path. If the path points to a single YAML file,
// it reads the policy directly from that file. If the path is a directory, it reads policies
// from all YAML files within the directory.
func ReadPolicies(policyPath string) ([]types.Policy, error) {
	// Check if the path is a file and if it's a YAML file
	if utils.IsItFile(policyPath) && utils.IsItYAML(policyPath) {
		// Read a single policy from the specified YAML file
		return readPolicyGen[[]types.Policy](policyPath)
	}

	// Read policies from all YAML files in the specified directory
	return readPoliciesFromDirectory(policyPath,
		func(path string) ([]types.Policy, error) {
			return readPolicyGen[[]types.Policy](path)
		})
}

// ReadPoliciesFromDirectory reads policies from all files in the specified directory
// using the provided policyFunc to handle each file.
func ReadPoliciesFromDirectory(policyPath string, policyFunc func(string) ([]types.Policy, error)) ([]types.Policy, error) {
	// Reuse the internal function to read policies from a directory
	return readPoliciesFromDirectory(policyPath, policyFunc)
}

// ReadPolicy reads a single policy from the given file path and returns it as type T.
func ReadPolicy[T any](path string) (T, error) {
	// Read a policy from the specified path using a generic read function
	return readPolicyGen[T](path)
}

// readPoliciesFromDirectory recursively walks through the specified directory,
// applying policyFunc to read each YAML file and collecting the resulting policies into a slice.
func readPoliciesFromDirectory(
	policyPath string,
	policyFunc func(string) ([]types.Policy, error),
) ([]types.Policy, error) {
	var policies []types.Policy

	// Walk the directory and process files
	err := filepath.Walk(policyPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			// Wrap the error with additional context about the path
			return fmt.Errorf("error accessing the path %q: %w", path, err)
		}

		// Check if the current item is a file and has a YAML extension
		if !info.IsDir() && utils.IsItYAML(path) {
			// Call the provided function to read the policy
			policy, err := policyFunc(path)
			if err != nil {
				// Wrap the error indicating which file failed to read
				return fmt.Errorf("could not read policy from file %q: %v", path, err)
			}

			// Append the retrieved policies to the slice
			policies = append(policies, policy...)
		}

		return nil
	})

	if err != nil {
		// Wrap the error indicating a problem occurred while walking the directory
		return policies, fmt.Errorf("error walking the path %q: %w", policyPath, err)
	}

	return policies, nil
}

// readPolicyGen reads the file at the given path and unmarshals its YAML content into the provided type T.
func readPolicyGen[T any](path string) (T, error) {
	var policy T

	// Read the file contents into memory
	file, err := os.ReadFile(path)
	if err != nil {
		// Wrap the error to provide context about the file that could not be read
		return policy, fmt.Errorf("could not read file %q: %v", path, err)
	}

	// Unmarshal the YAML file content into the specified type
	if err := yaml.Unmarshal(file, &policy); err != nil {
		// Wrap the unmarshalling error with context about the file
		return policy, fmt.Errorf("could not unmarshal YAML from file %q: %v", path, err)
	}

	return policy, nil
}
