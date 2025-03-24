package config

var (
	// DatabaseRolePrefix is the environment variable key for the database role prefix.
	DatabaseRolePrefix = "GM_DATABASE_ROLE_PREFIX"
	
	// DefaultDatabaseRolePrefix specifies the default prefix for database roles.
	DefaultDatabaseRolePrefix = "dwh_"
	
	// DatabaseType is the environment variable key for the type of database being used.
	DatabaseType = "GM_DATABASE_TYPE"
	
	// RuntestCleanup is the environment variable key that indicates whether to clean up the test environment after running tests.
	RuntestCleanup = "GM_RUNTEST_CLEANUP"
	
	// DefaultRuntestCleanup specifies the default behavior for test cleanup; "true" means cleanup will occur by default.
	DefaultRuntestCleanup = "true"
)
