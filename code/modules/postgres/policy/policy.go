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
	*base.Policy
	pool       *pgxpool.Pool // The connection pool used for database operations
	rolePrefix string
}

// New initializes a new PGPolicy instance with the provided pgxpool.Pool.
// This function takes a pointer to a pgxpool.Pool (PostgreSQL connection pool)
// as an argument and returns a pointer to a PGPolicy struct.
func New(ctx context.Context, config map[string]string) (*PGPolicy, error) {
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

	basePolicy := base.NewPolicy(config) // Initialize the base.Role

	return &PGPolicy{
		Policy: basePolicy,
		pool:   pool,
	}, nil
}
