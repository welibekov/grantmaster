package role

import (
	"context"
	"fmt"

	"github.com/welibekov/grantmaster/modules/config"
	"github.com/welibekov/grantmaster/modules/databaser"

	fgRole "github.com/welibekov/grantmaster/modules/fakegres/role"
	gpRole "github.com/welibekov/grantmaster/modules/greenplum/role"
	pgRole "github.com/welibekov/grantmaster/modules/postgres/role"
	"github.com/welibekov/grantmaster/modules/types"
)

func New(ctx context.Context, cfg map[string]string) (databaser.Roler, error) {
	switch types.DatabaseType(cfg[config.DatabaseType]) {
	case types.Postgres:
		// Initialize Postgres database
		return pgRole.New(ctx, cfg)
	case types.Fakegres:
		// Initialize Fakegres database
		return fgRole.New(cfg)
	// Initialize Greenplum database
	case types.Greenplum:
		return gpRole.New(ctx, cfg)
	default:
		// Return an error if the database type is not recognized
		return nil, fmt.Errorf("database type could not be found: GM_DATABASE_TYPE")
	}
}
