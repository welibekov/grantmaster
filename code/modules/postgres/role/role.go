package role

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/welibekov/grantmaster/modules/postgres/types"
	"github.com/welibekov/grantmaster/modules/role/base"
)

// PGRole represents a role in the database.
// It holds a reference to a connection pool to interact with the database.
type PGRole struct {
	*base.Role
	pool *pgxpool.Pool // The connection pool used for database interactions.
}

func New(ctx context.Context, config map[string]string) (*PGRole, error) {
	// Retrieve the connection string from the configuration map
	connString, found := config[types.ConnectionString]
	if !found {
		return nil, fmt.Errorf("%s not defined", types.ConnectionString) // Return an error if not found
	}

	// Create a new connection pool using the connection string
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, fmt.Errorf("couldn't connect to database: %v", err) // Return an error if the connection fails
	}

	defer func() {
		if err != nil { // close connection when error occured
			pool.Close()
		}
	}()

	baseRole := base.NewRole(config) // Initialize the base.Role

	// Return a new Postgres instance with the initialized database and connection string
	return &PGRole{
		Role: baseRole, // Set basic role
		pool: pool,     // Set initialized pool
	}, nil
}
