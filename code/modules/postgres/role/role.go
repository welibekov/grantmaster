package role

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

// New function creates a new PGRole instance given a pgxpool.Pool.
// It initializes the PGRole struct with the provided database connection pool.
func New(pool *pgxpool.Pool, rolePrefix string) *PGRole {
	return &PGRole{
		rolePrefix: rolePrefix,
		pool:       pool, // Assign the provided pool to the PGRole instance.
	}
}

// PGRole represents a role in the database.
// It holds a reference to a connection pool to interact with the database.
type PGRole struct {
	rolePrefix string        // rolePrefix is a prefix applied to role names.
	pool       *pgxpool.Pool // The connection pool used for database interactions.
}
