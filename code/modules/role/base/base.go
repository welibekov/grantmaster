package base

import (
	"context"
	"fmt"

	"github.com/welibekov/grantmaster/modules/role/types"
)

type Role struct {
	Prefix string
}

func NewRole(config map[string]string) *Role {
	return &Role{
		Prefix: config["GM_DATABASE_ROLE_PREFIX"], // Role filename prefix from the configuration.
	}
}

func (r *Role) Apply(context.Context, []types.Role) error {
	return r.notImplemented()
}

func (r *Role) Drop(context.Context, []types.Role) error {
	return r.notImplemented()
}

func (r *Role) Revoke(context.Context, []types.Role) error {
	return r.notImplemented()
}

func (r *Role) Grant(context.Context, []types.Role) error {
	return r.notImplemented()
}

func (r *Role) Get(context.Context) ([]types.Role, error) {
	return []types.Role{}, r.notImplemented()
}

func (r *Role) notImplemented() error {
	return fmt.Errorf("NYI")
}
