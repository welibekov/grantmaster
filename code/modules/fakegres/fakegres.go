package fakegres

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/welibekov/grantmaster/modules/database/base"
	"github.com/welibekov/grantmaster/modules/utils"
)

type Fakegres struct {
	*base.Database

	rootDir   string // Directory where fakegres data is stored
	roleDir   string
	policyDir string
}

// New creates a new instance of Fakegres with the provided configuration.
// If the GM_FAKEGRES_ROOTDIR is not specified in the config, it defaults to "/tmp/fakegres".
func New(config map[string]string) (*Fakegres, error) {
	// Retrieve the root directory from the configuration or set a default.
	rootDir, found := config["GM_FAKEGRES_ROOTDIR"]
	if !found {
		rootDir = "/tmp/fakegres"
	}

	fakegres := &Fakegres{
		roleDir:   filepath.Join(rootDir, "role"),
		policyDir: filepath.Join(rootDir, "policy"),
	}

	// Check if the specified root directory exists.
	for _, dir := range []string{fakegres.policyDir, fakegres.roleDir} {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			// If it doesn't exist, create the directory with default permissions.
			if err := os.MkdirAll(dir, os.ModePerm); err != nil {
				// Wrap and return the error for better context.
				return nil, fmt.Errorf("failed to create directory %s: %w", dir, err)
			}
		}
	}

	fakegres.rootDir = rootDir // Set the root directory
	fakegres.Database = base.NewDatabase()

	return fakegres, nil // Return the initialized Fakegres instance
}

// absPath constructs an absolute file path by joining the provided path components.
// If the resulting path does not have a ".yaml" extension, it appends ".yaml" to the end.
func (f *Fakegres) absPath(path ...string) string {
	// Join the provided path components into a single path string.
	result := filepath.Join(path...)

	// Check if the resulting path is not a YAML file.
	if !utils.IsItYAML(result) {
		// If it's not a YAML file, append the ".yaml" extension to the path.
		result += ".yaml"
	}

	// Return the final path, which will be a valid YAML file path.
	return result
}
