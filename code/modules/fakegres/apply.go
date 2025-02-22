package fakegres

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"

	"github.com/welibekov/grantmaster/modules/policy/types"
)

// apply is a method for the Fakegres struct that accepts a Policy and manages file storage based on that Policy.
func (f *Fakegres) applyPolicy(policy types.Policy) error {
	// Construct the filename based on the Username
	filename := filepath.Join(f.rootDir, (policy.Username + ".yaml"))

	updatePolicy, err := yaml.Marshal(&policy)
	if err != nil {
		return fmt.Errorf("failed to marshal policy: %w", err) // Return marshal error with context
	}

	// Try to read the existing file
	existingPolicy, err := ioutil.ReadFile(filename)
	if err != nil {
		// If file does not exist, create it
		if os.IsNotExist(err) {
			return ioutil.WriteFile(filename, updatePolicy, 0644)
		}

		return fmt.Errorf("failed to read file %s: %w", filename, err) // Return other errors with context
	}

	// Check if the roles are the same
	if bytes.Equal(existingPolicy, updatePolicy) {
		return nil // No change needed
	}

	// If roles are different, add the new roles to the file
	return ioutil.WriteFile(filename, updatePolicy, 0644)
}
