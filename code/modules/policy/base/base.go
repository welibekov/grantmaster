package base

import (
	"context"
	"fmt"

	"github.com/welibekov/grantmaster/modules/policy/types"
)

type Policy struct {
	RolePrefix string
}

func NewPolicy(config map[string]string) *Policy {
	return &Policy{
		RolePrefix: config["GM_DATABASE_ROLE_PREFIX"], // Policy filename prefix from the configuration.
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
