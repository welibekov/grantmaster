package fakegres

import (
	"context"
	"fmt"
	"os"

	"github.com/welibekov/grantmaster/modules/policy/types"
)

// ApplyPolicy processes a slice of policies and applies the specified actions (grant/revoke).
// It updates the policies map and ensures that any absent policies are removed.
func (f *Fakegres) ApplyPolicy(_ context.Context, policies []types.Policy) error {
	updatePolicesMap := make(map[string][]string) // Map to keep track of policy updates

	// Loop through each policy to apply it.
	for _, policy := range policies {
		updatePolicesMap[policy.Username] = policy.Roles // Update the map with the user's roles

		// Apply the individual policy.
		if err := apply(policy, f.absPath(f.policyDir, policy.Username)); err != nil {
			// Wrap and return the error with context about which policy failed.
			return fmt.Errorf("failed to apply policy for user %s: %w", policy.Username, err)
		}
	}

	// Remove any policies that are not present in the updated policies map.
	if err := f.removeAbsentPolicies(updatePolicesMap); err != nil {
		// Wrap and return the error for better clarity on failure.
		return fmt.Errorf("failed to remove absent policies: %w", err)
	}

	return nil // Successfully applied all policies
}

// removeAbsentPolicies removes policies for users that are not present in the updatePolicesMap.
// It compares the existing policies with the incoming updates and removes any policies
// for users that are no longer required.
func (f *Fakegres) removeAbsentPolicies(updatePolicesMap map[string][]string) error {
	// Read the existing policies from the storage.
	existingPoliciesMap, err := readExisting[types.Policy](f.policyDir,
		func(policies []types.Policy) map[string][]string {
			policiesMap := make(map[string][]string)
			for _, policy := range policies {
				policiesMap[policy.Username] = policy.Roles
			}

			return policiesMap
		},
	)
	if err != nil {
		// Wrap the error with additional context before returning.
		return fmt.Errorf("failed to read existing policies: %w", err)
	}

	// Iterate over existing policies to identify which ones need to be removed.
	for username := range existingPoliciesMap {
		_, found := updatePolicesMap[username]
		if !found {
			// If the username is not found in the update map, proceed to remove the policy.
			if err := os.Remove(f.absPath(f.policyDir, username)); err != nil {
				// Wrap the error with additional context before returning.
				return fmt.Errorf("failed to remove policy for user '%s': %w", username, err)
			}
		}
	}

	return nil
}
