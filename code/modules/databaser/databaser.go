package databaser

import "github.com/welibekov/grantmaster/modules/types"

// Databaser defines methods for managing policies and permissions.
type Databaser interface {
	// Apply applies a set of policies to the system.
	Apply([]types.Policy) error
}
