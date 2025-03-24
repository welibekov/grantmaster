package base

import (
	"context"
	"fmt"

	"github.com/welibekov/grantmaster/modules/config"
	"github.com/welibekov/grantmaster/modules/role/types"
)

// Role represents a role with an associated prefix.
type Role struct {
	Prefix string // Prefix for role filenames.
}

// NewRole creates a new Role instance using the provided configuration.
func NewRole(cfg map[string]string) *Role {
	return &Role{
		Prefix: cfg[config.DatabaseRolePrefix], // Role filename prefix from the configuration.
	}
}

// Apply applies the given roles, but is not yet implemented.
func (r *Role) Apply(ctx context.Context, roles []types.Role) error {
	return r.notImplemented()
}

// Drop drops the given roles, but is not yet implemented.
func (r *Role) Drop(ctx context.Context, roles []types.Role) error {
	return r.notImplemented()
}

// Revoke revokes the given roles, but is not yet implemented.
func (r *Role) Revoke(ctx context.Context, roles []types.Role) error {
	return r.notImplemented()
}

// Grant grants the given roles, but is not yet implemented.
func (r *Role) Grant(ctx context.Context, roles []types.Role) error {
	return r.notImplemented()
}

// Get retrieves the current roles, but is not yet implemented.
func (r *Role) Get(ctx context.Context) ([]types.Role, error) {
	return []types.Role{}, r.notImplemented()
}

// notImplemented returns an error indicating that the function is not yet implemented.
func (r *Role) notImplemented() error {
	return fmt.Errorf("NYI") // NYI stands for "Not Yet Implemented".
}
