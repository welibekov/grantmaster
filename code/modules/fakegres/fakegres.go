package fakegres

import (
	"context"
	"fmt"
	"os"

	"github.com/welibekov/grantmaster/modules/database/base"
	"github.com/welibekov/grantmaster/modules/types"
)

type Fakegres struct {
	*base.Database

	rootDir string // Directory where fakegres data is stored
}

// New creates a new instance of Fakegres with the provided configuration.
// If the GM_FAKEGRES_ROOTDIR is not specified in the config, it defaults to "/tmp/fakegres".
func New(config map[string]string) (*Fakegres, error) {
	fakegres := &Fakegres{}

	// Retrieve the root directory from the configuration or set a default.
	rootDir, found := config["GM_FAKEGRES_ROOTDIR"]
	if !found {
		rootDir = "/tmp/fakegres"
	}

	// Check if the specified root directory exists.
	if _, err := os.Stat(rootDir); os.IsNotExist(err) {
		// If it doesn't exist, create the directory with default permissions.
		if err := os.Mkdir(rootDir, os.ModePerm); err != nil {
			// Wrap and return the error for better context.
			return nil, fmt.Errorf("failed to create directory %s: %w", rootDir, err)
		}
	}

	fakegres.rootDir = rootDir // Set the root directory
	fakegres.Database = base.NewDatabase()

	return fakegres, nil // Return the initialized Fakegres instance
}

// Apply processes a slice of policies and applies the specified actions (grant/revoke).
// It updates the policies map and ensures that any absent policies are removed.
func (f *Fakegres) ApplyPolicy(_ context.Context, policies []types.Policy) error {
	updatePolicesMap := make(map[string][]string) // Map to keep track of policy updates

	// Loop through each policy to apply it.
	for _, policy := range policies {
		updatePolicesMap[policy.Username] = policy.Roles // Update the map with the user's roles

		// Apply the individual policy.
		if err := f.applyPolicy(policy); err != nil {
			// Wrap and return the error with context about which policy failed.
			return fmt.Errorf("failed to apply policy for user %s: %w", policy.Username, err)
		}
	}

	// Remove any policies that are not present in the updated policies map.
	if err := f.removeAbsentPolicies(updatePolicesMap); err != nil {
		// Wrap and return the error for better clarity on failure.
		return fmt.Errorf("failed to remove absent policies: %w", err)
	}

	return nil // Successfully applied all policies
}
