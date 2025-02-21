package base

import (
	"fmt"

	"github.com/welibekov/grantmaster/modules/types"
)

type Database struct{}

func NewDatabase() *Database {
	return &Database{}
}

func (d *Database) ApplyPolicy(policies []types.Policy) error {
	return d.notImplemented()
}

func (d *Database) ApplyRole(roles []types.Role) error {
	return d.notImplemented()
}

func (d *Database) notImplemented() error {
	return fmt.Errorf("NYI")
}
