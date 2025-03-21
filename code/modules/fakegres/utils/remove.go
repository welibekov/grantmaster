package utils

import (
	"os"
	"strings"
)

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
