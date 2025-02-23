package assets

import (
	"fmt"
)

// DetectDuplicated checks for duplicate items based on the keyExtractor().
// It returns an error if a duplicate is found.
func DetectDuplicated[T any](items []T, keyExtractor func(T) string) error {
	// Create a map to track items that have already been processed.
	keysMap := map[string]struct{}{}

	// Iterate over each item in the provided slice.
	for _, item := range items {
		key := keyExtractor(item)
		// Check if the key of the current item already exists in the map.
		_, found := keysMap[key]
		if found {
			// If the item is found, return an error indicating a duplicate.
			return fmt.Errorf("duplicated entry for '%s'", key)
		}

		// If the key is not found, add it to the map.
		keysMap[key] = struct{}{}
	}

	// If no duplicates were found, return nil indicating success.
	return nil
}
