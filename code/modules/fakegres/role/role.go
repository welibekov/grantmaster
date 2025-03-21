package role

import (
	"fmt"
	"path/filepath"

	"github.com/welibekov/grantmaster/modules/assets"
	"github.com/welibekov/grantmaster/modules/fakegres/utils"
	"github.com/welibekov/grantmaster/modules/role/base"
)

type FGRole struct {
	*base.Role

	rootDir string // Directory where fakegres data is stored
	roleDir string // Directory where role files are stored
}

func New(cfg map[string]string) (*FGRole, error) {
	// Initialize a FGPolicy instance with the appropriate directory paths.
	fgRole := &FGRole{
		Role: base.NewRole(cfg),

		// Retrieve the root directory from the configuration or set a default.
		rootDir: utils.GetRootDir(cfg),
		roleDir: filepath.Join(utils.GetRootDir(cfg), "role"), // Full path to the role directory.
	}

	// Check if the specified root directory exists and create necessary subdirectories.
	if err := assets.CreateDir(fgRole.roleDir); err != nil {
		return nil, fmt.Errorf("couldn't create directory %s: %v", fgRole.roleDir, err)
	}

	return fgRole, nil
}
