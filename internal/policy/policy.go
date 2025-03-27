package policy

import (
	"context"
	"fmt"

	"github.com/welibekov/grantmaster/internal/config"
	"github.com/welibekov/grantmaster/internal/databaser"
	fgPol "github.com/welibekov/grantmaster/internal/fakegres/policy"
	gpPol "github.com/welibekov/grantmaster/internal/greenplum/policy"
	pgPol "github.com/welibekov/grantmaster/internal/postgres/policy"
	"github.com/welibekov/grantmaster/internal/types"
)

// New creates a new Policier based on the specified database type in the configuration.
// It takes a context and a configuration map as input and returns a Policier interface or an error.
func New(ctx context.Context, cfg map[string]string) (databaser.Policier, error) {
	// Retrieve the database type from the configuration
	switch types.DatabaseType(cfg[config.DatabaseType]) {
	case types.Postgres:
		// Initialize and return a new Postgres Policier
		return pgPol.New(ctx, cfg)
	case types.Fakegres:
		// Initialize and return a new Fakegres Policier
		return fgPol.New(cfg)
	case types.Greenplum:
		// Initialize and return a new Greenplum Policier
		return gpPol.New(ctx, cfg)
	default:
		// Return an error if the database type is not recognized
		return nil, fmt.Errorf("database type could not be found: GM_DATABASE_TYPE")
	}
}
