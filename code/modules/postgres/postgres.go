package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"

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

	// Assign the newly created pool to the Postgres struct
	p.pool = pool

	// Retrieve existing policies from the database
	exisitingPolicies, err := pgPolicy.GetExisting(ctx, pool)
	if err != nil {
		return fmt.Errorf("couldn't apply policies: %v", err)
	}

	// Determine which policies need to be revoked based on the current and new policies
	revokePolicies := policy.WhatToRevoke(policies, exisitingPolicies)

	// Log the length of policies to be revoked for debugging purposes
	logrus.Debugln("Revoke policies length=", len(revokePolicies))

	// Revoke the identified policies from the database
	if err := p.revokePolicy(ctx, revokePolicies); err != nil {
		return fmt.Errorf("couldn't revoke policies: %v", err)
	}

	// Determine which policies need to be granted based on the current and new policies
	grantPolicies := policy.WhatToGrant(policies, exisitingPolicies)

	// Grant the new policies to the database
	return p.grantPolicy(ctx, grantPolicies)
}
