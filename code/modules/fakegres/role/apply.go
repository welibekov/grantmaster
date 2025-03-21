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

func (f *FGRole) Apply(_ context.Context, roles []types.Role) error {
	existingRoles, err := f.Get(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get exising roles: %v", err)
	}
	debug.OutputMarshal(existingRoles, "existing roles")

	revokeRoles := utils.Diff(roles, existingRoles)
	debug.OutputMarshal(revokeRoles, "revoke roles")
	debug.OutputMarshal(roles, "apply roles")

	// Remove revoked roles.
	err = fgUtils.Remove[types.Role](revokeRoles,
		func(item types.Role) string {
			return filepath.Join(f.roleDir, item.Name)
		})

	if err != nil {
		return fmt.Errorf("couldn't remove roles: %v", err)
	}

	err = fgUtils.Save(roles,
		func(item types.Role) string {
			return filepath.Join(f.roleDir, item.Name)
		})

	if err != nil {
		return fmt.Errorf("couldn't save roles: %v", err)
	}

	return nil
}
