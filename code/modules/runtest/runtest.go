package runtest

import (
	"fmt"

	"github.com/welibekov/grantmaster/modules/databaser"
	fgRuntest "github.com/welibekov/grantmaster/modules/fakegres/runtest"
	pgRuntest "github.com/welibekov/grantmaster/modules/postgres/runtest"
	"github.com/welibekov/grantmaster/modules/types"
)

func New(dbType types.DatabaseType, tests []string) (databaser.RunTesterer, error) {
	switch dbType {
	case types.Postgres:
		return pgRuntest.New(tests)
	case types.Fakegres:
		return fgRuntest.New(tests)
	default:
		return nil, fmt.Errorf("no database type %s found", dbType)
	}
}
