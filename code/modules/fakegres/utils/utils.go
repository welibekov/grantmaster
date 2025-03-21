package utils

// Define constants for the fakegres root directory configuration.
var (
	FakegresRootDir        = "GM_FAKEGRES_ROOTDIR" // Environment variable key for the root directory.
	DefaultFakegresRootDir = "/tmp/fakegres"        // Default root directory if none is specified in the configuration.
)

// GetRootDir retrieves the root directory from the provided configuration map.
// If the key for the root directory is not found, it returns a default value.
func GetRootDir(cfg map[string]string) string {
	// Retrieve the root directory from the configuration or set a default.
	rootDir, found := cfg[FakegresRootDir]
	if !found {
		rootDir = DefaultFakegresRootDir // Set to default value if not found.
	}

	return rootDir // Return the determined root directory.
}
