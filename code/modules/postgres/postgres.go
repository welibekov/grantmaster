package postgres

import (
	"fmt"

	"github.com/welibekov/grantmaster/modules/types"
)

type Postgres struct{}

func New(config map[string]string) (*Postgres, error) {
	return &Postgres{}, nil
}

func (p *Postgres) Apply([]types.Policy) error {
	return fmt.Errorf("NYI")
}

func (p *Postgres) Grant(string, []string) error {
	return fmt.Errorf("NYI")
}

func (p *Postgres) Revoke(string, []string) error {
	return fmt.Errorf("NYI")
}
