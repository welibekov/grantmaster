package fakegres

import (
	"fmt"

	"github.com/welibekov/grantmaster/modules/assets"
)

// readExisting reads all YAML files in the specified directory and unmarshals them into generic structs.
func readExisting[T any](rootDir string, toMap func(data []T) map[string][]string) (map[string][]string, error) {
	policies, err := assets.ReadAssetsFromDirectory(rootDir,
		func(path string) ([]T, error) {
			policies := []T{}

			policy, err := assets.ReadAsset[T](path)
			if err != nil {
				return policies, err
			}

			return append(policies, policy), nil
		})

	if err != nil {
		return nil, fmt.Errorf("error walking the path %q: %w", rootDir, err)
	}

	return toMap(policies), nil
}
