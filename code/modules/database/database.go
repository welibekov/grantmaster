package database

import (
	"fmt"

	"github.com/welibekov/grantmaster/modules/databaser"
	"github.com/welibekov/grantmaster/modules/fakegres"
	"github.com/welibekov/grantmaster/modules/postgres"
	"github.com/welibekov/grantmaster/modules/types"
)

// New creates a new Databaser instance based on the provided configuration.
func New(config types.Config) (databaser.Databaser, error) {
	switch types.DatabaseType(config.Database) {
	case types.Postgres:
		// Initialize Postgres database
		return postgres.New(config)
	case types.Fakegres:
		// Initialize Fakegres database
		return fakegres.New(config)
	default:
		// Return an error if the database type is not recognized
		return nil, fmt.Errorf("database type could not be found")
	}
}
