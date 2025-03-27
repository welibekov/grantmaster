package policy

import (
	"fmt"
	"path/filepath"

	"github.com/welibekov/grantmaster/internal/assets"
	"github.com/welibekov/grantmaster/internal/fakegres/utils"
	"github.com/welibekov/grantmaster/internal/policy/base"
)

// FGPolicy represents the policy specific to the Fakegres service.
// It embeds the base.Policy struct to inherit its fields and methods.
type FGPolicy struct {
	*base.Policy // Embedding the base policy for shared functionality

	rootDir   string // Directory where fakegres data is stored
	policyDir string // Directory where policy files are stored
}

// New initializes a new FGPolicy instance with the provided configuration.
// It sets up the necessary directories and checks for their existence.
func New(cfg map[string]string) (*FGPolicy, error) {
	// Create a new instance of FGPolicy and initialize its fields.
	fgPol := &FGPolicy{
		Policy: base.NewPolicy(cfg), // Create a new base policy using the configuration

		// Retrieve the root directory from the configuration or set a default.
		rootDir:   utils.GetRootDir(cfg), // Get the root directory from the config
		policyDir: filepath.Join(utils.GetRootDir(cfg), "policy"), // Full path to the policy directory.
	}

	// Check if the specified policy directory exists; create it if it doesn't.
	if err := assets.CreateDir(fgPol.policyDir); err != nil {
		// Return an error if directory creation fails, including the directory path and the error.
		return nil, fmt.Errorf("couldn't create directory %s: %v", fgPol.policyDir, err)
	}

	// Return the new FGPolicy instance if no errors occurred.
	return fgPol, nil
}
