package policy

import (
	"context"
	"fmt"

	"github.com/welibekov/grantmaster/modules/config"
	"github.com/welibekov/grantmaster/modules/databaser"
	fgPol "github.com/welibekov/grantmaster/modules/fakegres/policy"
	gpPol "github.com/welibekov/grantmaster/modules/greenplum/policy"
	pgPol "github.com/welibekov/grantmaster/modules/postgres/policy"
	"github.com/welibekov/grantmaster/modules/types"
)

func New(ctx context.Context, cfg map[string]string) (databaser.Policier, error) {
	switch types.DatabaseType(cfg[config.DatabaseType]) {
	case types.Postgres:
		// Initialize Postgres database
		return pgPol.New(ctx, cfg)
	case types.Fakegres:
		// Initialize Fakegres database
		return fgPol.New(cfg)
	// Initialize Greenplum database
	case types.Greenplum:
		return gpPol.New(ctx, cfg)
	default:
		// Return an error if the database type is not recognized
		return nil, fmt.Errorf("database type could not be found: GM_DATABASE_TYPE")
	}
}
