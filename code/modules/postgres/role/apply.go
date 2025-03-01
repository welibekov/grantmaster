package role

import (
	"context"

	"github.com/welibekov/grantmaster/modules/role"
	"github.com/welibekov/grantmaster/modules/role/types"
	"github.com/welibekov/grantmaster/modules/utils/debug"
)

func (p *PGRole) Apply(ctx context.Context, roles []types.Role) error {
	roles = p.addRolePrefix(roles)
	debug.OutputMarshal(roles, "roles requested")

	existingRoles, err := p.GetExisting(ctx)
	if err != nil {
		return err
	}
	debug.OutputMarshal(existingRoles, "exising roles") // Log the generated query for debugging purposes

	createRoles := role.WhatToCreate(roles, existingRoles)
	debug.OutputMarshal(createRoles, "roles to be created") // Log the generated query for debugging purposes

	removeRoles := role.WhatToRemove(roles, existingRoles)
	debug.OutputMarshal(removeRoles, "roles to be removed") // Log the generated query for debugging purposes

	grantRoles := role.Diff(roles, existingRoles)
	revokeRoles := role.Diff(existingRoles, roles)

	debug.OutputMarshal(grantRoles, "roles to grant")
	debug.OutputMarshal(revokeRoles, "roles to revoke")

	if len(grantRoles) > 0 {
		if err := p.Grant(ctx, grantRoles); err != nil {
			return err
		}
	}

	if len(removeRoles) > 0 {
		if err := p.Drop(ctx, removeRoles); err != nil {
			return err
		}
	}

	// temporary disabled
	//if len(grantRoles) > 0 {
	//	if err := p.Grant(ctx, grantRoles); err != nil {
	//		return err
	//	}
	//}

	return nil
}
