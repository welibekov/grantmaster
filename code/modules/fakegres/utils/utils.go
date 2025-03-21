package utils

var (
	FakegresRootDir        = "GM_FAKEGRES_ROOTDIR"
	DefaultFakegresRootDir = "/tmp/fakegres"
)

func GetRootDir(cfg map[string]string) string {
	// Retrieve the root directory from the configuration or set a default.
	rootDir, found := cfg[FakegresRootDir]
	if !found {
		rootDir = DefaultFakegresRootDir // Default value for the root directory.
	}

	return rootDir
}
