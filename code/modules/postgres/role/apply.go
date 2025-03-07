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

	//existingRoles, err := p.GetExisting(ctx)
	existingRoles, err := p.Get(ctx)
	if err != nil {
		return err
	}
	debug.OutputMarshal(existingRoles, "exising roles") // Log the generated query for debugging purposes

	// FIXME: This two functions are not relevant now,
	// but we still need some mechanism for drop roles.
	//createRoles := role.WhatToCreate(roles, existingRoles)
	//debug.OutputMarshal(createRoles, "roles to be created") // Log the generated query for debugging purposes

	//removeRoles := role.WhatToRemove(roles, existingRoles)
	//debug.OutputMarshal(removeRoles, "roles to be removed") // Log the generated query for debugging purposes

	grantRoles := role.Diff(roles, existingRoles)
	revokeRoles := role.Diff(existingRoles, roles)
	dropRoles := role.WhatToRemove(roles, existingRoles)

	debug.OutputMarshal(grantRoles, "roles to grant")
	debug.OutputMarshal(revokeRoles, "roles to revoke")

	if len(grantRoles) > 0 {
		if err := p.Grant(ctx, grantRoles); err != nil {
			return err
		}
	}

	if len(revokeRoles) > 0 {
		if err := p.Revoke(ctx, revokeRoles); err != nil {
			return err
		}
	}

	if len(dropRoles) > 0 {
		if err := p.Drop(ctx, dropRoles); err != nil {
			return err
		}
	}

	return nil
}
