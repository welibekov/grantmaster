package cmd

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/welibekov/grantmaster/internal/config"
)

// init initializes the rootCmd flags.
func init() {
	// Define the loglevel flag with a default value of "info".
	rootCmd.PersistentFlags().String("loglevel", "info", "Set loglevel")
}

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "gm", // The command name.
	Short: "Grant Master CLI tool", // A brief description of the command.
	Long:  `A command line tool to apply database permissions based on policies.`, // A detailed description of the command.
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// Get the loglevel from the command flags.
		level, err := cmd.Flags().GetString("loglevel")
		if err != nil {
			return err // Return an error if unable to get loglevel.
		}

		// Set the GM_DATABASE_TYPE environment variable based on the caller command.
		if len(os.Args) > 2 {
			os.Setenv(config.DatabaseType, os.Args[1])
		}

		// Set the log level as specified.
		return setLogLevel(level)
	},
}

// setLogLevel sets the logging level for logrus based on the provided level string.
func setLogLevel(level string) error {
	// Parse the log level from the string representation.
	parsed, err := logrus.ParseLevel(level)
	if err != nil {
		return fmt.Errorf("invalid loglevel %s: %v", level, err) // Return an error if the log level is invalid.
	}

	// Set the log level for logrus.
	logrus.SetLevel(parsed)

	return nil // Return nil if the log level is successfully set.
}

// Execute runs the root command, handling any errors that occur.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err) // Print the error to stdout.
		os.Exit(1)       // Exit the program with a non-zero status code.
	}
}
