package config

import (
	"os"
	"strings"
)

// Load retrieves environment variables that start with the "GM_" prefix
// and organizes them into a map of key-value pairs. If no specific
// role prefix for the database is found in the environment variables,
// it sets a default role prefix.
func Load() map[string]string {
	// Create a map to hold the configuration key-value pairs
	config := make(map[string]string)

	// Iterate over all environment variables
	for _, env := range os.Environ() {
		// Split each environment variable into key and value
		kv := strings.SplitN(env, "=", 2)

		// Check if the split resulted in a key and value, and if the key
		// starts with the "GM_" prefix
		if len(kv) == 2 && strings.HasPrefix(kv[0], "GM_") {
			// Add the key-value pair to the config map
			config[kv[0]] = kv[1]
		}
	}

	// Check if the database role prefix exists in the config
	_, found := config[DatabaseRolePrefix]
	// If not found, set it to a default value
	if !found {
		config[DatabaseRolePrefix] = DefaultDatabaseRolePrefix
	}

	_, found = config[RuntestCleanup]
	if !found {
		config[RuntestCleanup] = DefaultRuntestCleanup
	}

	// Return the populated config map
	return config
}
