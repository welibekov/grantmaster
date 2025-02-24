package policy

import "github.com/jackc/pgx/v5/pgxpool"

// New initializes a new PGPolicy instance with the provided pgxpool.Pool.
// This function takes a pointer to a pgxpool.Pool (PostgreSQL connection pool)
// as an argument and returns a pointer to a PGPolicy struct.
func New(pool *pgxpool.Pool) *PGPolicy {
	return &PGPolicy{
		pool: pool, // Store the provided pool in the PGPolicy instance
	}
}

// PGPolicy represents a policy implementation using a PostgreSQL connection pool.
// It contains a single field 'pool' which is a reference to the connection pool.
type PGPolicy struct {
	pool *pgxpool.Pool // The connection pool used for database operations
}
