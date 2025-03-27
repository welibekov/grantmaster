package databaser

import (
	"context"

	"github.com/welibekov/grantmaster/internal/policy/types"
)

// Policier defines the interface for managing policies in the system.
type Policier interface {
	// Apply applies a set of policies and returns an error if the operation fails.
	Apply(ctx context.Context, policies []types.Policy) error
	
	// Grant grants a set of policies and returns an error if the operation fails.
	Grant(ctx context.Context, policies []types.Policy) error
	
	// Revoke revokes a set of policies and returns an error if the operation fails.
	Revoke(ctx context.Context, policies []types.Policy) error
	
	// Get retrieves the current policies and returns them along with an error if the operation fails.
	Get(ctx context.Context) ([]types.Policy, error)
}
