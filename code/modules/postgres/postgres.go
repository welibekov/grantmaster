package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"gopkg.in/yaml.v3"

	"github.com/welibekov/grantmaster/modules/database/base"
	"github.com/welibekov/grantmaster/modules/policy"
	"github.com/welibekov/grantmaster/modules/policy/types"
	pgPolicy "github.com/welibekov/grantmaster/modules/postgres/policy"
)

type Postgres struct {
	*base.Database               // Embedding base.Database to inherit its properties and methods
	pool           *pgxpool.Pool // connection pool for managing database connections
	connString     string        // connection string for establishing a connection to the database
}

func New(config map[string]string) (*Postgres, error) {
	// Retrieve the connection string from the configuration map
	connString, found := config["GM_POSTGRES_CONN_STRING"]
	if !found {
		return nil, fmt.Errorf("GM_POSTGRES_CONN_STRING not defined") // Return an error if not found
	}

	// Return a new Postgres instance with the initialized database and connection string
	return &Postgres{
		Database:   base.NewDatabase(), // Initialize the base.Database
		connString: connString,         // Set the connection string
	}, nil

}

// ApplyPolicy method applies a set of policies to the Postgres database.
func (p *Postgres) ApplyPolicy(ctx context.Context, policies []types.Policy) error {
	// Create a new connection pool using the connection string
	pool, err := pgxpool.New(ctx, p.connString)
	if err != nil {
		return fmt.Errorf("couldn't connect to database: %v", err) // Return an error if the connection fails
	}
	defer pool.Close() // Ensure that the connection pool is closed when the function exits

	exisitingPolicies, err := pgPolicy.GetExisting(ctx, pool)
	if err != nil {
		return fmt.Errorf("couldn't apply policies: %v", err)
	}

	applyPolicy := policy.Compare(policies, exisitingPolicies)

	yamlBytes, err := yaml.Marshal(applyPolicy)
	if err != nil {
		return fmt.Errorf("couldn't marshal yaml: %v")
	}

	os.Stdout.Write(yamlBytes)

	// This indicates that the implementation of the ApplyPolicy method is not yet complete
	return fmt.Errorf("NYI") // NYI: Not Yet Implemented
}
