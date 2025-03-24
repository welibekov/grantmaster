package policy

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/welibekov/grantmaster/modules/policy/base"
	"github.com/welibekov/grantmaster/modules/postgres/types"
)

// PGPolicy represents a policy implementation using a PostgreSQL connection pool.
// It contains a single field 'pool' which is a reference to the connection pool.
type PGPolicy struct {
	*base.Policy // Embedding the base.Policy to leverage its functionality
	pool         *pgxpool.Pool // The connection pool used for database operations
	rolePrefix   string         // Prefix for roles, if applicable
}

// New initializes a new PGPolicy instance with the provided pgxpool.Pool.
// This function takes a context and a configuration map as arguments and
// returns a pointer to a PGPolicy struct along with any error encountered.
func New(ctx context.Context, config map[string]string) (*PGPolicy, error) {
	// Retrieve the connection string from the configuration map
	connString, found := config[types.ConnectionString]
	if !found {
		// Return an error if the connection string is not found in the configuration
		return nil, fmt.Errorf("%s not defined", types.ConnectionString)
	}

	// Create a new connection pool using the connection string
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		// Return an error if the connection to the database fails
		return nil, fmt.Errorf("couldn't connect to database: %v", err)
	}

	// Defer a function to close the connection pool if an error occurs
	defer func() {
		if err != nil { // Close connection if an error occurred during initialization
			pool.Close()
		}
	}()

	// Initialize the base.Policy with the provided configuration
	basePolicy := base.NewPolicy(config)

	// Return a new PGPolicy instance while including the initialized base.Policy and connection pool
	return &PGPolicy{
		Policy: basePolicy, // Set the embedded Policy
		pool:   pool,       // Set the connection pool
	}, nil
}
