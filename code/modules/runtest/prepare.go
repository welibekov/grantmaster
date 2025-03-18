package runtest

import (
	"fmt"

	pgEnv "github.com/welibekov/grantmaster/modules/postgres/env"
	"github.com/welibekov/grantmaster/modules/types"
)

func (r *Runtest) prepareTestEnv(execDir string) (func() error, error) {
	switch r.Type {
	case types.Postgres:
		return pgEnv.Prepare(execDir)
	default:
		return nil, fmt.Errorf("couldn't prepare env for %s, not found", r.Type.ToString())
	}
}
