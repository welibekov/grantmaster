package fakegres

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/welibekov/grantmaster/modules/database/base"
	"github.com/welibekov/grantmaster/modules/utils"
)

// Fakegres represents a mock database structure that stores roles and policies.
type Fakegres struct {
	*base.Database // Embedding the base database type for database functionality.

	rootDir    string // Directory where fakegres data is stored
	roleDir    string // Directory where role files are stored
	policyDir  string // Directory where policy files are stored
	rolePrefix string // Prefix to be applied to role filenames
}

// New creates a new instance of Fakegres with the provided configuration.
// If the GM_FAKEGRES_ROOTDIR is not specified in the config, it defaults to "/tmp/fakegres".
func New(config map[string]string) (*Fakegres, error) {
	// Retrieve the root directory from the configuration or set a default.
	rootDir, found := config["GM_FAKEGRES_ROOTDIR"]
	if !found {
		rootDir = "/tmp/fakegres" // Default value for the root directory.
	}

	// Initialize a Fakegres instance with the appropriate directory paths.
	fakegres := &Fakegres{
		roleDir:    filepath.Join(rootDir, "role"),    // Full path to the role directory.
		policyDir:  filepath.Join(rootDir, "policy"),  // Full path to the policy directory.
		rolePrefix: config["GM_DATABASE_ROLE_PREFIX"],  // Role filename prefix from the configuration.
	}

	// Check if the specified root directory exists and create necessary subdirectories.
	for _, dir := range []string{fakegres.policyDir, fakegres.roleDir} {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			// If the directory does not exist, attempt to create it.
			if err := os.MkdirAll(dir, os.ModePerm); err != nil {
				// Wrap and return the error for better context.
				return nil, fmt.Errorf("failed to create directory %s: %w", dir, err)
			}
		}
	}

	fakegres.rootDir = rootDir // Set the root directory for the Fakegres instance.
	fakegres.Database = base.NewDatabase() // Initialize the embedded base database.

	return fakegres, nil // Return the initialized Fakegres instance.
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

// absPathRole constructs the absolute path for a role file.
// It uses a role filename prefix and ensures the path points to a YAML file.
func (f *Fakegres) absPathRole(path ...string) string {
	// Generate the absolute path for the role file.
	absPath := f.absPath(path...)

	// Extract the filename from the absolute path.
	filename := filepath.Base(absPath)
	// Get the directory of the absolute path.
	directory := filepath.Dir(absPath)

	// Combine the directory and the filename prefixed with rolePrefix to form the final path.
	return filepath.Join(directory, f.rolePrefix+filename)
}
