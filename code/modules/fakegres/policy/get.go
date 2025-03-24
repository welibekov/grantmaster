package policy

import (
	"context"
	"fmt"

	"github.com/welibekov/grantmaster/modules/assets"
	"github.com/welibekov/grantmaster/modules/policy/types"
)

// Get retrieves the existing policies from the specified directory.
// It reads all policy assets, unmarshals them, and returns them as a slice of types.Policy.
// In case of an error during reading or unmarshalling, an error is returned with additional context.
func (p *FGPolicy) Get(_ context.Context) ([]types.Policy, error) {
	// Read the existing policies from the storage.
	policies, err := assets.ReadAssetsFromDirectory[types.Policy](p.policyDir,
		func(path string) ([]types.Policy, error) {
			// Initialize a local slice to hold the policies read from a single asset.
			policies := []types.Policy{}

			// Read a single asset (policy) from the specified path.
			policy, err := assets.ReadAsset[types.Policy](path)
			if err != nil {
				// Return the empty slice and the error if reading the asset fails.
				return policies, err
			}

			// Append the successfully read policy to the slice.
			// The append function doesn't modify the original slice, so we return a new slice containing the old policies and the new policy.
			return append(policies, policy), nil
		})

	if err != nil {
		// Wrap the error with additional context before returning.
		return []types.Policy{}, fmt.Errorf("failed to read existing policies: %w", err)
	}

	// Return the slice of policies read from the storage.
	// At this point, policies contain all successfully read policies from the assets.
	return policies, nil
}
