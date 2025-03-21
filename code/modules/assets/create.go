package assets

import (
	"fmt"
	"os"
)

func CreateDir(path string) error {
	// Check if the specified root directory exists and create necessary subdirectories.
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// If the directory does not exist, attempt to create it.
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			// Wrap and return the error for better context.
			fmt.Errorf("failed to create directory %s: %w", path, err)
		}
	}

	return nil
}
