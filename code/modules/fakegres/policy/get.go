package policy

import (
	"context"
	"fmt"

	"github.com/welibekov/grantmaster/modules/assets"
	"github.com/welibekov/grantmaster/modules/policy/types"
)

func (p *FGPolicy) Get(_ context.Context) ([]types.Policy, error) {
	// Read the existing policies from the storage.
	policies, err := assets.ReadAssetsFromDirectory[types.Policy](p.policyDir,
		func(path string) ([]types.Policy, error) {
			policies := []types.Policy{}

			policy, err := assets.ReadAsset[types.Policy](path)
			if err != nil {
				return policies, err
			}

			return append(policies, policy), nil
		})

	if err != nil {
		// Wrap the error with additional context before returning.
		return []types.Policy{}, fmt.Errorf("failed to read existing policies: %w", err)
	}

	return policies, nil
}
