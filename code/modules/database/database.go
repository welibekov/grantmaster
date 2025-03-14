package database

import (
	"fmt"

	"github.com/welibekov/grantmaster/modules/databaser"
	"github.com/welibekov/grantmaster/modules/fakegres"
	"github.com/welibekov/grantmaster/modules/greenplum"
	"github.com/welibekov/grantmaster/modules/types"
)

// New creates a new Databaser instance based on the provided configuration.
func New(config map[string]string) (databaser.Databaser, error) {
	switch types.DatabaseType(config["GM_DATABASE_TYPE"]) {
	case types.Fakegres:
		// Initialize Fakegres database
		return fakegres.New(config)
		// Initialize Greenplum database
	case types.Greenplum:
		return greenplum.New(config)
	default:
		// Return an error if the database type is not recognized
		return nil, fmt.Errorf("database type could not be found: GM_DATABASE_TYPE")
	}
}
