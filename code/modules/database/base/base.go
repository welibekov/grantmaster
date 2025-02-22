package base

import (
	"context"
	"fmt"

	"github.com/welibekov/grantmaster/modules/policy/types"
)

type Database struct{}

func NewDatabase() *Database {
	return &Database{}
}

func (d *Database) ApplyPolicy(context.Context, []types.Policy) error {
	return d.notImplemented()
}

func (d *Database) ApplyRole(context.Context, []types.Role) error {
	return d.notImplemented()
}

func (d *Database) notImplemented() error {
	return fmt.Errorf("NYI")
}
