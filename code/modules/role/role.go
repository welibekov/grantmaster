package role

import (
	"context"
	"fmt"

	"github.com/welibekov/grantmaster/modules/databaser"
	//fakegres "github.com/welibekov/grantmaster/modules/fakegres/role"
	//greenplum "github.com/welibekov/grantmaster/modules/greenplum/role"
	pgRole "github.com/welibekov/grantmaster/modules/postgres/role"
	"github.com/welibekov/grantmaster/modules/types"
)

func New(ctx context.Context, config map[string]string) (databaser.Roler, error) {
	switch types.DatabaseType(config["GM_DATABASE_TYPE"]) {
	case types.Postgres:
		// Initialize Postgres database
		return pgRole.New(ctx, config)
	//case types.Fakegres:
	// Initialize Fakegres database
	//	return fakegres.New(config)
	// Initialize Greenplum database
	//case types.Greenplum:
	//	return greenplum.New(config)
	default:
		// Return an error if the database type is not recognized
		return nil, fmt.Errorf("database type could not be found: GM_DATABASE_TYPE")
	}
}
