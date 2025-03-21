package policy

import (
	"fmt"
	"path/filepath"

	"github.com/welibekov/grantmaster/modules/assets"
	"github.com/welibekov/grantmaster/modules/fakegres/utils"
	"github.com/welibekov/grantmaster/modules/policy/base"
)

type FGPolicy struct {
	*base.Policy

	rootDir   string // Directory where fakegres data is stored
	policyDir string // Directory where policy files are stored
}

func New(cfg map[string]string) (*FGPolicy, error) {
	// Initialize a FGPolicy instance with the appropriate directory paths.
	fgPol := &FGPolicy{
		Policy: base.NewPolicy(cfg),

		// Retrieve the root directory from the configuration or set a default.
		rootDir:   utils.GetRootDir(cfg),
		policyDir: filepath.Join(utils.GetRootDir(cfg), "policy"), // Full path to the policy directory.
	}

	// Check if the specified root directory exists and create necessary subdirectories.
	if err := assets.CreateDir(fgPol.policyDir); err != nil {
		return nil, fmt.Errorf("couldn't create directory %s: %v", fgPol.policyDir, err)
	}

	return fgPol, nil
}
