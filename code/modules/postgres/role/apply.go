package role

import (
	"context"

	"github.com/welibekov/grantmaster/modules/role/types"
	"github.com/welibekov/grantmaster/modules/role/utils"
	"github.com/welibekov/grantmaster/modules/utils/debug"
)

func (p *PGRole) Apply(ctx context.Context, roles []types.Role) error {
	defer p.pool.Close() // Ensure that the connection pool is closed when the function exits

	debug.OutputMarshal(roles, "roles requested")

	existingRoles, err := p.Get(ctx)
	if err != nil {
		return err
	}
	debug.OutputMarshal(existingRoles, "exising roles") // Log the generated query for debugging purposes

	grantRoles := utils.Diff(roles, existingRoles)
	revokeRoles := utils.Diff(existingRoles, roles)
	dropRoles := utils.WhatToDrop(roles, existingRoles)

	debug.OutputMarshal(grantRoles, "roles to grant")
	debug.OutputMarshal(revokeRoles, "roles to revoke")
	debug.OutputMarshal(revokeRoles, "roles to drop")

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
