package base

import (
	"context"
	"fmt"

	"github.com/welibekov/grantmaster/modules/config"
	"github.com/welibekov/grantmaster/modules/policy/types"
)

type Policy struct {
	RolePrefix string
}

func NewPolicy(cfg map[string]string) *Policy {
	return &Policy{
		RolePrefix: cfg[config.DatabaseRolePrefix], // Policy filename prefix from the configuration.
	}
}

func (r *Policy) Apply(context.Context, []types.Policy) error {
	return r.notImplemented()
}

func (r *Policy) Revoke(context.Context, []types.Policy) error {
	return r.notImplemented()
}

func (r *Policy) Grant(context.Context, []types.Policy) error {
	return r.notImplemented()
}

func (r *Policy) Get(context.Context) ([]types.Policy, error) {
	return []types.Policy{}, r.notImplemented()
}

func (r *Policy) notImplemented() error {
	return fmt.Errorf("NYI")
}
