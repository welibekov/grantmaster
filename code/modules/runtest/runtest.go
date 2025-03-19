package runtest

import (
	"fmt"

	"github.com/welibekov/grantmaster/modules/databaser"
	pgRuntest "github.com/welibekov/grantmaster/modules/postgres/runtest"
	"github.com/welibekov/grantmaster/modules/types"
)

func New(config map[string]string, tests []string) (databaser.RunTesterer, error) {
	switch dbType := types.DatabaseType(config["GM_DATABASE_TYPE"]); dbType {
	case types.Postgres:
		return pgRuntest.New(tests)
	default:
		return nil, fmt.Errorf("no database type %s found", dbType)
	}
}
