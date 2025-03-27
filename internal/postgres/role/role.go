package role

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/welibekov/grantmaster/internal/postgres/types"
	"github.com/welibekov/grantmaster/internal/role/base"
)

// PGRole represents a role in the database.
// It holds a reference to a connection pool to interact with the database.
type PGRole struct {
	*base.Role     // Embedding the base Role type for role-related functionality.
	pool *pgxpool.Pool // The connection pool used for database interactions.
}

// New creates a new PGRole instance using the given context and configuration.
// It initializes the database connection pool and the base Role.
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

	// Ensure the connection pool is closed if an error occurs
	defer func() {
		if err != nil { // close connection when error occurred
			pool.Close()
		}
	}()

	baseRole := base.NewRole(config) // Initialize the base.Role

	// Return a new PGRole instance with the initialized base role and connection pool
	return &PGRole{
		Role: baseRole, // Set base role initialized
		pool: pool,     // Set the initialized connection pool
	}, nil
}
