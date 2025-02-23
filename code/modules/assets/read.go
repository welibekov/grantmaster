package assets

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/welibekov/grantmaster/modules/utils"
	"gopkg.in/yaml.v3"
)

// ReadAssets reads assets from the specified path. If the path is a file and is
// a YAML file, it reads the asset directly. Otherwise, it attempts to read assets
// from all YAML files within a directory.
func ReadAssets[T any](assetPath string) ([]T, error) {
	if utils.IsItFile(assetPath) && utils.IsItYAML(assetPath) {
		// Read a single asset from a YAML file
		return readAsset[[]T](assetPath)
	}

	// Read assets from all YAML files in the specified directory
	return readAssetFromDirectory(assetPath,
		func(path string) ([]T, error) {
			return readAsset[[]T](path)
		})
}

// ReadAssetsFromDirectory reads assets from all files in a directory using
// the provided assetFunc to handle each file.
func ReadAssetsFromDirectory[T any](assetPath string, assetFunc func(string) ([]T, error)) ([]T, error) {
	return readAssetFromDirectory[T](assetPath, assetFunc)
}

// ReadPolicy reads a single asset from the given file path and returns it as type T.
func ReadAsset[T any](path string) (T, error) {
	return readAsset[T](path)
}

// readFromDirectory walks through the specified directory, applying assetFunc to read
// each YAML file and collecting the resulting assets.
func readAssetFromDirectory[T any](
	assetPath string,
	assetFunc func(string) ([]T, error),
) ([]T, error) {
	var assets []T

	err := filepath.Walk(assetPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			// Wrap the error with additional context
			return fmt.Errorf("error accessing the path %q: %w", path, err)
		}

		// Check if the file has a .yaml or .yml extension
		if !info.IsDir() && utils.IsItYAML(path) {
			asset, err := assetFunc(path)
			if err != nil {
				// Wrap the error with context indicating which file failed
				return fmt.Errorf("could not read asset from file %q: %v", path, err)
			}

			// Add the retrieved assets to the slice
			assets = append(assets, asset...)
		}

		return nil
	})

	if err != nil {
		// Wrap the error indicating there was a problem walking the directory
		return assets, fmt.Errorf("error walking the path %q: %w", assetPath, err)
	}

	return assets, nil
}

// readAsset reads the file at the given path and unmarshals its YAML content into the provided type T.
func readAsset[T any](path string) (T, error) {
	var data T

	// Read the file contents
	file, err := os.ReadFile(path)
	if err != nil {
		// Wrap the error to provide context
		return data, fmt.Errorf("could not read file %q: %v", path, err)
	}

	// Unmarshal the YAML file content into the specified type
	if err := yaml.Unmarshal(file, &data); err != nil {
		// Wrap the unmarshalling error with context
		return data, fmt.Errorf("could not unmarshal YAML from file %q: %v", path, err)
	}

	return data, nil
}
