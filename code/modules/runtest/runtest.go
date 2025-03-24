package runtest

import (
	"fmt"

	"github.com/welibekov/grantmaster/modules/databaser"
	fgRuntest "github.com/welibekov/grantmaster/modules/fakegres/runtest"
	pgRuntest "github.com/welibekov/grantmaster/modules/postgres/runtest"
	"github.com/welibekov/grantmaster/modules/types"
)

// New creates a new instance of RunTesterer based on the provided database type.
// It initializes and returns the appropriate testing implementation for the specified database.
// If the dbType is not recognized, it returns an error indicating the database type is not supported.
//
// Currently, this function only supports two database types: Postgres and Fakegres. 
// If additional database types are required in the future, they should be added to this 
// switch statement along with their corresponding initialization logic.
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
