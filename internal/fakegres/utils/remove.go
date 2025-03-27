package utils

import (
	"os"
	"strings"
)

// Remove deletes files whose paths are provided by the getPath function for each item in the items slice.
// It appends ".yaml" to the path if it does not already start with ".yaml".
//
// Parameters:
// - items: a slice of items of any type T.
// - getPath: a function that returns the file path as a string for each item in the items slice.
//
// Returns:
// - An error if any file cannot be removed or if a stat error occurs.
func Remove[T any](items []T, getPath func(item T) string) error {
	for _, item := range items {
		// Retrieve the file path from the current item using the provided getPath function.
		path := getPath(item)

		// Check if the path does not already start with ".yaml".
		if !strings.HasPrefix(path, ".yaml") {
			// Append ".yaml" to the path if it doesn't start with it.
			path += ".yaml"
		}

		// Check if the file exists by attempting to get its stats.
		if _, err := os.Stat(path); err == nil {
			// If the file exists, attempt to remove it.
			if err := os.Remove(path); err != nil {
				// Return any error encountered while trying to remove the file.
				return err
			}
		}
	}

	// Return nil if all files are successfully processed without errors.
	return nil
}
