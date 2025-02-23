package databaser

import (
	"context"

	polTypes "github.com/welibekov/grantmaster/modules/policy/types"
	rolTypes "github.com/welibekov/grantmaster/modules/role/types"
)

// Databaser defines methods for managing policies and permissions.
type Databaser interface {
	// Apply applies a set of policies to the system.
	ApplyPolicy(context.Context, []polTypes.Policy) error

	// ApplyRoles applices a set of roles to the system.
	ApplyRole(context.Context, []rolTypes.Role) error
}
