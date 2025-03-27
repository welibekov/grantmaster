package role

import (
	"context"
	"fmt"
	"path/filepath"

	fgUtils "github.com/welibekov/grantmaster/internal/fakegres/utils"
	"github.com/welibekov/grantmaster/internal/role/types"
	"github.com/welibekov/grantmaster/internal/role/utils"
	"github.com/welibekov/grantmaster/internal/utils/debug"
)

// Apply updates the roles by removing revoked roles and saving the new set of roles.
// It takes a context and a slice of roles to apply.
func (f *FGRole) Apply(_ context.Context, roles []types.Role) error {
	// Retrieve existing roles from the storage.
	existingRoles, err := f.Get(context.Background())
	if err != nil {
		// Return an error if retrieval of existing roles fails.
		return fmt.Errorf("couldn't get existing roles: %v", err)
	}
	debug.OutputMarshal(existingRoles, "existing roles")

	// Determine which roles need to be revoked (those not present in the new roles).
	revokeRoles := utils.Diff(roles, existingRoles)
	debug.OutputMarshal(revokeRoles, "revoke roles")
	debug.OutputMarshal(roles, "apply roles")

	// Remove revoked roles from the storage using a helper function.
	err = fgUtils.Remove[types.Role](revokeRoles,
		func(item types.Role) string {
			// Construct the file path for the role to be removed from storage.
			return filepath.Join(f.roleDir, item.Name)
		})

	if err != nil {
		// Return an error if removal of roles fails.
		return fmt.Errorf("couldn't remove roles: %v", err)
	}

	// Save the updated list of roles to the storage using a helper function.
	err = fgUtils.Save(roles,
		func(item types.Role) string {
			// Construct the file path for the role to be saved in storage.
			return filepath.Join(f.roleDir, item.Name)
		})

	if err != nil {
		// Return an error if saving of roles fails.
		return fmt.Errorf("couldn't save roles: %v", err)
	}

	// Successfully applied the roles; return nil indicating no error occurred.
	return nil
}
