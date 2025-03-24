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
		// Extract the key for the current item using the provided keyExtractor function.
		key := keyExtractor(item)

		// Check if the current item's key already exists in the keysMap.
		_, found := keysMap[key]
		if found {
			// If the key is found, return an error indicating a duplicate entry.
			return fmt.Errorf("duplicated entry for '%s'", key)
		}

		// If the key is not found, add it to keysMap to track its occurrence.
		keysMap[key] = struct{}{}
	}

	// If no duplicates were found after processing all items, return nil indicating success.
	return nil
}
