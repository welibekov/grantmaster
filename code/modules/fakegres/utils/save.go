package utils

import (
	"fmt"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v3"
)

func Save[T any](items []T, getPath func(item T) string) error {
	if len(items) > 0 {
		for _, item := range items {
			path := getPath(item)

			if !strings.HasPrefix(path, ".yaml") {
				path += ".yaml"
			}

			yamlBytes, err := yaml.Marshal(item)
			if err != nil {
				return fmt.Errorf("failed to marshal item: %w", err) // Return marshal error with context
			}

			if err := ioutil.WriteFile(path, yamlBytes, 0644); err != nil {
				return fmt.Errorf("couldn't save data for %s: %v", path, err)
			}
		}
	}

	return nil
}
