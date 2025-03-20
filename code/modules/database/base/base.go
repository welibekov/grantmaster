package base

import (
	"context"
	"fmt"

	"github.com/welibekov/grantmaster/modules/config"
	polTypes "github.com/welibekov/grantmaster/modules/policy/types"
	rolTypes "github.com/welibekov/grantmaster/modules/role/types"
)

type Database struct {
	RolePrefix string
}

func NewDatabase(cfg map[string]string) *Database {
	return &Database{
		RolePrefix: cfg[config.DatabaseRolePrefix], // Role filename prefix from the configuration.
	}
}

func (d *Database) ApplyPolicy(context.Context, []polTypes.Policy) error {
	return d.notImplemented()
}

func (d *Database) ApplyRole(context.Context, []rolTypes.Role) error {
	return d.notImplemented()
}

func (d *Database) notImplemented() error {
	return fmt.Errorf("NYI")
}
