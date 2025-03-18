package types

// DatabaseType represents the type of database being used.
type DatabaseType string

var (
	Postgres  DatabaseType = "postgres"  // Postgres is a constant for the PostgreSQL database type.
	Greenplum DatabaseType = "greenplum" // Greenplum is a constant for the Greenplum database type.
	Fakegres  DatabaseType = "fakegres"  // Fakegres is a constant for a mocked or fake database type.
)

func (d DatabaseType) ToString() string {
	return string(d)
}
