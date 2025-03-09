package greenplum

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"                        // Importing the pgxpool package for PostgreSQL connection pooling
	"github.com/welibekov/grantmaster/modules/database/base" // Importing base database module
	"github.com/welibekov/grantmaster/modules/policy/types"  // Importing types module for policy definitions
)

// Greenplum struct represents a connection to a Greenplum database.
type Greenplum struct {
	*base.Database               // Embedding base.Database to inherit its properties and methods
	pool           *pgxpool.Pool // connection pool for managing database connections
	connString     string        // connection string for establishing a connection to the database
}

// New function initializes a new Greenplum instance using the provided configuration map.
func New(config map[string]string) (*Greenplum, error) {
	// Retrieve the connection string from the configuration map
	connString, found := config["GM_GREENPLUM_CONN_STRING"]
	if !found {
		return nil, fmt.Errorf("GM_GREENPLUM_CONN_STRING not defined") // Return an error if not found
	}

	// Return a new Greenplum instance with the initialized database and connection string
	return &Greenplum{
		Database:   base.NewDatabase(config), // Initialize the base.Database
		connString: connString,               // Set the connection string
	}, nil
}

// ApplyPolicy method applies a set of policies to the Greenplum database.
func (g *Greenplum) ApplyPolicy(ctx context.Context, policies []types.Policy) error {
	// Create a new connection pool using the connection string
	pool, err := pgxpool.New(ctx, g.connString)
	if err != nil {
		return fmt.Errorf("couldn't connect to database: %v", err) // Return an error if the connection fails
	}
	defer pool.Close() // Ensure that the connection pool is closed when the function exits

	// Assign the newly created connection pool to the Greenplum struct
	g.pool = pool

	// This indicates that the implementation of the ApplyPolicy method is not yet complete
	return fmt.Errorf("NYI") // NYI: Not Yet Implemented
}
