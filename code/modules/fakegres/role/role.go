package role

import (
	"fmt"
	"path/filepath"

	"github.com/welibekov/grantmaster/modules/assets"
	"github.com/welibekov/grantmaster/modules/fakegres/utils"
	"github.com/welibekov/grantmaster/modules/role/base"
)

// FGRole represents a role in the fakegres system with specific directory configurations.
type FGRole struct {
	*base.Role    // Embedding base Role structure to leverage its functionality.
	rootDir string // Directory where fakegres data is stored.
	roleDir string // Directory where role files are stored.
}

// New initializes a new FGRole instance based on the provided configuration map.
// It also sets up the necessary directories for storing role and fakegres data.
func New(cfg map[string]string) (*FGRole, error) {
	// Initialize a FGRole instance with the appropriate directory paths.
	fgRole := &FGRole{
		Role: base.NewRole(cfg), // Create a new base Role using the provided configuration.

		// Retrieve the root directory from the configuration or set a default.
		rootDir: utils.GetRootDir(cfg),
		// Full path to the role directory under the root directory.
		roleDir: filepath.Join(utils.GetRootDir(cfg), "role"),
	}

	// Check if the specified role directory exists; if not, create it.
	if err := assets.CreateDir(fgRole.roleDir); err != nil {
		// Return an error if the directory creation fails, including the path and error message.
		return nil, fmt.Errorf("couldn't create directory %s: %v", fgRole.roleDir, err)
	}

	// Return the newly created FGRole instance.
	return fgRole, nil
}
