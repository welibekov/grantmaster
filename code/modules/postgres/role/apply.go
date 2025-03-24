package role

import (
	"context"

	"github.com/welibekov/grantmaster/modules/role/types"
	"github.com/welibekov/grantmaster/modules/role/utils"
	"github.com/welibekov/grantmaster/modules/utils/debug"
)

// Apply updates the roles in the database by granting, revoking, and dropping roles 
// based on the difference between the provided roles and existing roles.
//
// It closes the connection pool when exiting the function, ensuring 
// that resources are properly released.
func (p *PGRole) Apply(ctx context.Context, roles []types.Role) error {
	defer p.pool.Close() // Ensure that the connection pool is closed when the function exits.

	debug.OutputMarshal(roles, "roles requested") // Log the requested roles for debugging.

	// Retrieve existing roles from the database.
	existingRoles, err := p.Get(ctx)
	if err != nil {
		return err // Return error if fetching existing roles fails.
	}
	debug.OutputMarshal(existingRoles, "existing roles") // Log existing roles for comparison.

	// Calculate the differences between the requested roles and existing roles.
	grantRoles := utils.Diff(roles, existingRoles)
	revokeRoles := utils.Diff(existingRoles, roles)
	dropRoles := utils.WhatToDrop(roles, existingRoles)

	// Log the roles that will be granted, revoked, and dropped for debugging.
	debug.OutputMarshal(grantRoles, "roles to grant")
	debug.OutputMarshal(revokeRoles, "roles to revoke")
	debug.OutputMarshal(dropRoles, "roles to drop")

	// Grant any roles that are new.
	if len(grantRoles) > 0 {
		if err := p.Grant(ctx, grantRoles); err != nil {
			return err // Return error if granting roles fails.
		}
	}

	// Revoke any roles that are no longer needed.
	if len(revokeRoles) > 0 {
		if err := p.Revoke(ctx, revokeRoles); err != nil {
			return err // Return error if revoking roles fails.
		}
	}

	// Drop any roles that should be removed from the system.
	if len(dropRoles) > 0 {
		if err := p.Drop(ctx, dropRoles); err != nil {
			return err // Return error if dropping roles fails.
		}
	}

	return nil // Return nil if all operations succeed without errors.
}
