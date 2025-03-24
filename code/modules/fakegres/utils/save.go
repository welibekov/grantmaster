package utils

import (
	"fmt"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v3"
)

// Save serializes a list of items of any type to YAML and saves each item to a file.
// The file path for each item is determined by the provided getPath function.
// If the generated path does not have a ".yaml" suffix, ".yaml" is appended to it.
func Save[T any](items []T, getPath func(item T) string) error {
	// Check if there are any items to save
	if len(items) > 0 {
		// Iterate over each item in the provided slice
		for _, item := range items {
			// Get the file path based on the item using the provided getPath function
			path := getPath(item)

			// Ensure the file path has a .yaml extension
			if !strings.HasSuffix(path, ".yaml") {
				path += ".yaml"
			}

			// Serialize the current item to YAML format
			yamlBytes, err := yaml.Marshal(item)
			if err != nil {
				// Return an error if marshalling fails, wrapped with contextual information
				return fmt.Errorf("failed to marshal item: %w", err)
			}

			// Write the serialized YAML bytes to a file at the specified path
			if err := ioutil.WriteFile(path, yamlBytes, 0644); err != nil {
				// Return an error if writing the file fails, with contextual information
				return fmt.Errorf("couldn't save data for %s: %v", path, err)
			}
		}
	}

	// Return nil if all items were saved successfully
	return nil
}
