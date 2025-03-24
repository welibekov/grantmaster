package types

// DatabaseType represents the type of database being used.
type DatabaseType string

var (
	// Postgres is a constant for the PostgreSQL database type.
	Postgres DatabaseType = "postgres"  
	
	// Greenplum is a constant for the Greenplum database type.
	Greenplum DatabaseType = "greenplum" 
	
	// Fakegres is a constant for a mocked or fake database type.
	Fakegres DatabaseType = "fakegres"  
)

// Databases is a slice containing all supported DatabaseType constants.
var Databases = []DatabaseType{Postgres, Greenplum, Fakegres}

// ToString returns the string representation of the DatabaseType.
func (d DatabaseType) ToString() string {
	return string(d)
}
