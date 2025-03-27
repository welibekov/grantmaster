package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/welibekov/grantmaster/internal/config"
	"github.com/welibekov/grantmaster/internal/runtest"
	"github.com/welibekov/grantmaster/internal/types"

	"github.com/spf13/cobra"
)

// init initializes the command by adding the gmRuntestCmd to the provided commands.
func init() {
	for _, gmCmd := range []*cobra.Command{gmPostgresCmd, gmFakegresCmd, gmGreenplumCmd} {
		gmCmd.AddCommand(gmRuntestCmd)
	}
}

// gmRuntestCmd represents the runtest command which is used to run tests based on the database type.
var gmRuntestCmd = &cobra.Command{
	Use:   "runtest",
	Short: "Run tests of specific test type (default in GM_DATABASE_TYPE)",
	// RunE is the execution function for the command.
	RunE: func(cmd *cobra.Command, args []string) error {
		// Load configuration settings
		cfg := config.Load()

		// Determine the database type from the configuration
		dbType := types.DatabaseType(cfg[config.DatabaseType])

		// If no arguments are provided, default to the tests directory for the specified dbType
		if len(args) == 0 {
			args = append(args, filepath.Join("tests", dbType.ToString()))
		}

		// Slice to hold test file paths
		tests := []string{}
		for _, arg := range args {
			// Walk through the directory provided in the argument
			err := filepath.WalkDir(arg, func(path string, d os.DirEntry, err error) error {
				if err != nil {
					return err // Return error if encountered while walking the directory
				}

				// Check if the entry is a file and has a .bash suffix
				if !d.IsDir() && strings.HasSuffix(d.Name(), ".bash") {
					tests = append(tests, path) // Append test path to tests slice
				}

				return nil
			})

			if err != nil {
				fmt.Println("Error:", err) // Print error if walking directory fails
			}
		}

		// If no tests were found, return an error
		if len(tests) == 0 {
			return fmt.Errorf("no tests files found")
		}

		// Parse the cleanup flag from the configuration
		cleanup, err := strconv.ParseBool(cfg[config.RuntestCleanup])
		if err != nil {
			return fmt.Errorf("Wrong value for %s: %v", cfg[config.RuntestCleanup], err)
		}

		// Create a new runtest instance with the specified database type and tests
		rt, err := runtest.New(dbType, tests)
		if err != nil {
			return err // Return error if runtest creation fails
		}

		// Prepare the runtest environment
		cleanupFn, err := rt.Prepare()
		if err != nil {
			return fmt.Errorf("couldn't prepare runtest env: %v", err) // Return error if preparation fails
		}

		// If cleanup is enabled, defer the cleanup function
		if cleanup {
			defer cleanupFn()
		}

		// Execute the tests and return any errors that occur
		return rt.Execute()
	},
}
