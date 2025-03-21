package role

import (
	"context"
	"fmt"

	"github.com/welibekov/grantmaster/modules/assets"
	"github.com/welibekov/grantmaster/modules/role/types"
)

func (p *FGRole) Get(_ context.Context) ([]types.Role, error) {
	// Read the existing roles from the storage.
	roles, err := assets.ReadAssetsFromDirectory[types.Role](p.roleDir,
		func(path string) ([]types.Role, error) {
			roles := []types.Role{}

			role, err := assets.ReadAsset[types.Role](path)
			if err != nil {
				return roles, err
			}

			return append(roles, role), nil
		})

	if err != nil {
		// Wrap the error with additional context before returning.
		return []types.Role{}, fmt.Errorf("failed to read existing roles: %w", err)
	}

	return roles, nil
}
