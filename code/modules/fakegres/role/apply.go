package role

import (
	"context"
	"fmt"
	"path/filepath"

	fgUtils "github.com/welibekov/grantmaster/modules/fakegres/utils"
	"github.com/welibekov/grantmaster/modules/role/types"
	"github.com/welibekov/grantmaster/modules/role/utils"
	"github.com/welibekov/grantmaster/modules/utils/debug"
)

// Apply updates the roles by removing revoked roles and saving the new set of roles.
// It takes a context and a slice of roles to apply.
func (f *FGRole) Apply(_ context.Context, roles []types.Role) error {
	// Retrieve existing roles from the storage.
	existingRoles, err := f.Get(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get existing roles: %v", err)
	}
	debug.OutputMarshal(existingRoles, "existing roles")

	// Determine which roles need to be revoked (those not present in the new roles).
	revokeRoles := utils.Diff(roles, existingRoles)
	debug.OutputMarshal(revokeRoles, "revoke roles")
	debug.OutputMarshal(roles, "apply roles")

	// Remove revoked roles from the storage.
	err = fgUtils.Remove[types.Role](revokeRoles,
		func(item types.Role) string {
			return filepath.Join(f.roleDir, item.Name)
		})

	if err != nil {
		return fmt.Errorf("couldn't remove roles: %v", err)
	}

	// Save the updated list of roles to the storage.
	err = fgUtils.Save(roles,
		func(item types.Role) string {
			return filepath.Join(f.roleDir, item.Name)
		})

	if err != nil {
		return fmt.Errorf("couldn't save roles: %v", err)
	}

	// Successfully applied the roles.
	return nil
}
