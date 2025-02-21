package postgres

import (
	"fmt"

	"github.com/welibekov/grantmaster/modules/database/base"
	"github.com/welibekov/grantmaster/modules/types"
)

type Postgres struct {
	*base.Database
}

func New(config map[string]string) (*Postgres, error) {
	return &Postgres{
		base.NewDatabase(),
	}, nil

}

func (p *Postgres) Apply([]types.Policy) error {
	return fmt.Errorf("NYI")
}
