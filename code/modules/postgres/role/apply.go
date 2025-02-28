package role

import (
	"context"

	"github.com/welibekov/grantmaster/modules/role"
	"github.com/welibekov/grantmaster/modules/role/types"
	"github.com/welibekov/grantmaster/modules/utils/debug"
)

func (p *PGRole) Apply(ctx context.Context, roles []types.Role) error {
	roles = p.addRolePrefix(roles)

	existingRoles, err := p.GetExisting(ctx)
	if err != nil {
		return err
	}

	debug.OutputMarshal(existingRoles, "exising roles") // Log the generated query for debugging purposes

	createRoles := role.WhatToCreate(roles, existingRoles)
	removeRoles := role.WhatToRemove(roles, existingRoles)

	debug.OutputMarshal(createRoles, "roles to be created") // Log the generated query for debugging purposes
	debug.OutputMarshal(removeRoles, "roles to be removed") // Log the generated query for debugging purposes

	if len(createRoles) > 0 {
		if err := p.Create(ctx, createRoles); err != nil {
			return err
		}
	}

	if len(removeRoles) > 0 {
		if err := p.Drop(ctx, removeRoles); err != nil {
			return err
		}
	}

	return nil
}
