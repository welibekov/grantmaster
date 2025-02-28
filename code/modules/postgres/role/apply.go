package role

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/welibekov/grantmaster/modules/role"
	"github.com/welibekov/grantmaster/modules/role/types"
)

func (p *PGRole) Apply(ctx context.Context, roles []types.Role) error {
	roles = p.addRolePrefix(roles)

	existingRoles, err := p.GetExisting(ctx)
	if err != nil {
		return err
	}

	logrus.Debugln(existingRoles) // Log the generated query for debugging purposes

	createRoles := role.WhatToCreate(roles, existingRoles)
	removeRoles := role.WhatToRemove(roles, existingRoles)

	logrus.Debugln(createRoles) // Log the generated query for debugging purposes
	logrus.Debugln(removeRoles) // Log the generated query for debugging purposes

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
