package role

import (
	"context"
	"fmt"

	"github.com/welibekov/grantmaster/modules/assets"
	"github.com/welibekov/grantmaster/modules/role/types"
)

// Get retrieves the list of roles from the storage.
func (p *FGRole) Get(_ context.Context) ([]types.Role, error) {
	// Read the existing roles from the specified directory.
	roles, err := assets.ReadAssetsFromDirectory[types.Role](p.roleDir,
		func(path string) ([]types.Role, error) {
			roles := []types.Role{}

			// Read an individual role asset from the given path.
			role, err := assets.ReadAsset[types.Role](path)
			if err != nil {
				// Return the slice of roles and the error if reading the role fails.
				return roles, err
			}

			// Append the read role to the roles slice and return it.
			return append(roles, role), nil
		})

	if err != nil {
		// Wrap the error with additional context before returning.
		return []types.Role{}, fmt.Errorf("failed to read existing roles: %w", err)
	}

	// Return the retrieved roles and no error.
	return roles, nil
}
