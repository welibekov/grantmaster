package base

import (
	"context"
	"fmt"

	polTypes "github.com/welibekov/grantmaster/modules/policy/types"
	rolTypes "github.com/welibekov/grantmaster/modules/role/types"
)

type Database struct{}

func NewDatabase() *Database {
	return &Database{}
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
