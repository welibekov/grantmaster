package assets

import (
	"fmt"
	"os"
)

// CreateDir checks if the specified directory exists, and if not, it creates it along with any necessary parent directories.
func CreateDir(path string) error {
	// Check if the specified directory exists.
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// If the directory does not exist, attempt to create it and its parents.
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			// Wrap and return the error for better context.
			return fmt.Errorf("failed to create directory %s: %w", path, err)
		}
	}

	// Return nil if the directory already exists or if it was successfully created.
	return nil
}
