package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/welibekov/grantmaster/modules/database/base"
	polTypes "github.com/welibekov/grantmaster/modules/policy/types"
	pgPolicy "github.com/welibekov/grantmaster/modules/postgres/policy"
	pgRole "github.com/welibekov/grantmaster/modules/postgres/role"
	rolTypes "github.com/welibekov/grantmaster/modules/role/types"
)

type Postgres struct {
	*base.Database               // Embedding base.Database to inherit its properties and methods
	pool           *pgxpool.Pool // connection pool for managing database connections
	connString     string        // connection string for establishing a connection to the database
}

func New(config map[string]string) (*Postgres, error) {
	// Retrieve the connection string from the configuration map
	connString, found := config[pgConnectionString]
	if !found {
		return nil, fmt.Errorf("%s not defined", pgConnectionString) // Return an error if not found
	}

	// Return a new Postgres instance with the initialized database and connection string
	return &Postgres{
		Database:   base.NewDatabase(config), // Initialize the base.Database
		connString: connString,               // Set the connection string
	}, nil

}

// ApplyPolicy method applies a set of policies to the Postgres database.
func (p *Postgres) ApplyPolicy(ctx context.Context, policies []polTypes.Policy) error {
	// Create a new connection pool using the connection string
	pool, err := pgxpool.New(ctx, p.connString)
	if err != nil {
		return fmt.Errorf("couldn't connect to database: %v", err) // Return an error if the connection fails
	}
	defer pool.Close() // Ensure that the connection pool is closed when the function exits

	// Assign the newly created pool to the PGPolicy struct
	pgpol := pgPolicy.New(pool, p.RolePrefix)

	// Apply policies to Postgres database
	return pgpol.Apply(ctx, policies)
}

// ApplyRole method creates a set of roles to the Postgres database
func (p *Postgres) ApplyRole(ctx context.Context, roles []rolTypes.Role) error {
	// Create a new connection pool using the connection string
	pool, err := pgxpool.New(ctx, p.connString)
	if err != nil {
		return fmt.Errorf("couldn't connect to database: %v", err) // Return an error if the connection fails
	}
	defer pool.Close() // Ensure that the connection pool is closed when the function exits

	// Assign the newly created pool to the PGRole struct
	pgrol := pgRole.New(pool, p.RolePrefix)

	// Apply policies to Postgres database
	return pgrol.Apply(ctx, roles)
}
