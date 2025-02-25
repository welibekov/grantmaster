package role

import (
	"context"

	"github.com/welibekov/grantmaster/modules/role"
	"github.com/welibekov/grantmaster/modules/role/types"
)

func (p *PGRole) Apply(ctx context.Context, roles []types.Role) error {
	roles = p.addRolePrefix(roles)

	existingRoles, err := p.GetExisting(ctx)
	if err != nil {
		return err
	}

	createRoles := role.WhatToCreate(roles, existingRoles)
	removeRoles := role.WhatToRemove(roles, existingRoles)

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
