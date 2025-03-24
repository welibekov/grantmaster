package utils

// Define constants for the fakegres root directory configuration.
var (
	FakegresRootDir        = "GM_FAKEGRES_ROOTDIR" // Environment variable key for the root directory.
	DefaultFakegresRootDir = "/tmp/fakegres"        // Default root directory if none is specified in the configuration.
)

// GetRootDir retrieves the root directory from the provided configuration map.
// If the key for the root directory is not found, it returns a default value.
func GetRootDir(cfg map[string]string) string {
	// Attempt to retrieve the root directory from the configuration map.
	rootDir, found := cfg[FakegresRootDir]
	if !found {
		rootDir = DefaultFakegresRootDir // Set to the default value if the key is not present in the configuration.
	}

	return rootDir // Return the determined root directory.
}
