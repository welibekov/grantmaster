package databaser

import (
	"context"

	"github.com/welibekov/grantmaster/modules/role/types"
)

// Roler is an interface that defines methods for managing roles
// in a system. It allows for applying, dropping, revoking, and 
// granting roles, as well as retrieving the current set of roles.
type Roler interface {
	// Apply adds new roles to the system or updates existing roles.
	Apply(ctx context.Context, roles []types.Role) error
	
	// Drop removes roles from the system.
	Drop(ctx context.Context, roles []types.Role) error
	
	// Revoke removes permissions associated with specified roles.
	Revoke(ctx context.Context, roles []types.Role) error
	
	// Grant assigns permissions to the specified roles.
	Grant(ctx context.Context, roles []types.Role) error
	
	// Get retrieves the currently assigned roles in the system.
	Get(ctx context.Context) ([]types.Role, error)
}
