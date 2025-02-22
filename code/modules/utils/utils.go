package utils

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

// IsItFile checks if the given path corresponds to a file and not a directory.
func IsItFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		logrus.Warnln(err)
		return false
	}

	return !info.IsDir()
}

// IsItYAML checks if the given file path has a .yaml or .yml extension, indicating it's a YAML file.
func IsItYAML(path string) bool {
	return filepath.Ext(path) == ".yaml" || filepath.Ext(path) == ".yml"
}
