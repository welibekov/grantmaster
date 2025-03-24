package utils

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

// IsItFile checks if the given path corresponds to a file and not a directory.
// It returns true if the path is a file, and false if it is a directory or an error occurs.
func IsItFile(path string) bool {
	info, err := os.Stat(path) // Get the file information
	if err != nil {
		logrus.Warnln(err) // Log a warning if there is an error getting the file info
		return false // Return false if an error occurs (e.g., path does not exist)
	}

	return !info.IsDir() // Return true if it is not a directory (meaning it is a file)
}

// IsItYAML checks if the given file path has a .yaml or .yml extension,
// indicating that it's a YAML file. It returns true for valid YAML extensions.
func IsItYAML(path string) bool {
	// Check if the file extension is .yaml or .yml
	return filepath.Ext(path) == ".yaml" || filepath.Ext(path) == ".yml"
}
