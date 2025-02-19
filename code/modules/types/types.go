package types

// DatabaseType represents the type of database being used.
type DatabaseType string

var (
	Postgres DatabaseType // Postgres is a constant for the PostgreSQL database type.
	Fakegres DatabaseType // Fakegres is a constant for a mocked or fake database type.
)
