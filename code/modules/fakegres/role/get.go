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
			// Initialize a slice to hold the roles.
			roles := []types.Role{}

			// Read an individual role asset from the given path.
			role, err := assets.ReadAsset[types.Role](path)
			if err != nil {
				// Return an empty slice of roles and the error if reading the role fails.
				return roles, err
			}

			// Append the read role to the roles slice and return the updated slice.
			return append(roles, role), nil
		})

	if err != nil {
		// Wrap the error with additional context before returning to provide more information.
		return []types.Role{}, fmt.Errorf("failed to read existing roles: %w", err)
	}

	// Return the retrieved roles and indicate that there was no error.
	return roles, nil
}
