package utils

import (
	"os"
	"strings"
)

// Remove deletes files whose paths are provided by the getPath function for each item in the items slice.
// It appends ".yaml" to the path if it does not already start with it.
// 
// Parameters:
// - items: a slice of items of any type T.
// - getPath: a function that returns the file path as a string for each item in the items slice.
//
// Returns:
// - An error if any file cannot be removed or if a stat error occurs.
func Remove[T any](items []T, getPath func(item T) string) error {
	for _, item := range items {
		path := getPath(item)

		if !strings.HasPrefix(path, ".yaml") {
			path += ".yaml"
		}

		if _, err := os.Stat(path); err == nil {
			if err := os.Remove(path); err != nil {
				return err
			}
		}
	}

	return nil
}
