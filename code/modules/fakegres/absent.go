package fakegres

import (
	"fmt"
	"os"
	"path/filepath"
)

// removeAbsentPolicies removes policies for users that are not present in the updatePolicesMap.
// It compares the existing policies with the incoming updates and removes any policies
// for users that are no longer required.
func (f *Fakegres) removeAbsentPolicies(updatePolicesMap map[string][]string) error {
	// Read the existing policies from the storage.
	existingPoliciesMap, err := f.readExistingPolicies()
	if err != nil {
		// Wrap the error with additional context before returning.
		return fmt.Errorf("failed to read existing policies: %w", err)
	}

	// Iterate over existing policies to identify which ones need to be removed.
	for username := range existingPoliciesMap {
		_, found := updatePolicesMap[username]
		if !found {
			// If the username is not found in the update map, proceed to remove the policy.
			if err := f.removePolicy(username); err != nil {
				// Wrap the error with additional context before returning.
				return fmt.Errorf("failed to remove policy for user '%s': %w", username, err)
			}
		}
	}

	return nil
}

// removePolicy removes the policy file associated with the given username.
// It constructs the file path using the root directory and the username.
func (f *Fakegres) removePolicy(username string) error {
	// Construct the path to the policy file for the user.
	filePath := filepath.Join(f.rootDir, (username + ".yaml"))

	// Attempt to remove the policy file and return any error encountered.
	if err := os.Remove(filePath); err != nil {
		return fmt.Errorf("failed to delete policy file '%s': %w", filePath, err)
	}

	return nil
}
