package fakegres

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

// apply is a functions for the Fakegres struct that accepts a T and manages file storage based on that T.
func apply(data interface{}, filename string) error {
	updateData, err := yaml.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err) // Return marshal error with context
	}

	// Try to read the existing file
	existingData, err := ioutil.ReadFile(filename)
	if err != nil {
		// If file does not exist, create it
		if os.IsNotExist(err) {
			return ioutil.WriteFile(filename, updateData, 0644)
		}

		return fmt.Errorf("failed to read file %s: %w", filename, err) // Return other errors with context
	}

	// Check if the roles are the same
	if bytes.Equal(existingData, updateData) {
		return nil // No change needed
	}

	// If roles are different, add the new roles to the file
	return ioutil.WriteFile(filename, updateData, 0644)
}
