package utils

import (
	"fmt"

	"github.com/welibekov/grantmaster/modules/types"
)

// DetectDuplicated checks for duplicate policies based on the username.
// It returns an error if a duplicate is found.
func DetectDuplicated(policies []types.Policy) error {
	// Create a map to track usernames that have already been processed.
	policiesMap := map[string]struct{}{}

	// Iterate over each policy in the provided slice.
	for _, policy := range policies {
		// Check if the username of the current policy already exists in the map.
		_, found := policiesMap[policy.Username]
		if found {
			// If the username is found, return an error indicating a duplicate.
			return fmt.Errorf("duplicated policy for '%s' username", policy.Username)
		}

		// If the username is not found, add it to the map.
		policiesMap[policy.Username] = struct{}{}
	}

	// If no duplicates were found, return nil indicating success.
	return nil
}
