package base

import (
	"context"
	"fmt"

	"github.com/welibekov/grantmaster/modules/config"
	"github.com/welibekov/grantmaster/modules/policy/types"
)

// Policy represents a policy with a role prefix.
type Policy struct {
	RolePrefix string // Role prefix for the policy.
}

// NewPolicy creates a new Policy instance using the provided configuration map.
// It initializes the RolePrefix based on the value from the configuration.
func NewPolicy(cfg map[string]string) *Policy {
	return &Policy{
		RolePrefix: cfg[config.DatabaseRolePrefix], // Policy filename prefix from the configuration.
	}
}

// Apply applies a set of policies. Currently not implemented and returns an error indicating so.
func (r *Policy) Apply(ctx context.Context, policies []types.Policy) error {
	return r.notImplemented()
}

// Revoke revokes a set of policies. Currently not implemented and returns an error indicating so.
func (r *Policy) Revoke(ctx context.Context, policies []types.Policy) error {
	return r.notImplemented()
}

// Grant grants a set of policies. Currently not implemented and returns an error indicating so.
func (r *Policy) Grant(ctx context.Context, policies []types.Policy) error {
	return r.notImplemented()
}

// Get retrieves the current set of policies. Currently not implemented and returns an empty slice and an error indicating so.
func (r *Policy) Get(ctx context.Context) ([]types.Policy, error) {
	return []types.Policy{}, r.notImplemented()
}

// notImplemented returns an error signaling that the method has not yet been implemented.
func (r *Policy) notImplemented() error {
	return fmt.Errorf("NYI") // NYI stands for "Not Yet Implemented".
}
