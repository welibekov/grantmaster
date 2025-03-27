package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// init function adds the database commands to the root command.
func init() {
	for _, gmCmd := range []*cobra.Command{
		gmPostgresCmd, gmFakegresCmd, gmGreenplumCmd,
	} {
		rootCmd.AddCommand(gmCmd) // Add each database command to the root command
	}
}

// gmPostgresCmd represents the command for Postgres database.
var gmPostgresCmd = &cobra.Command{
	Use:   "postgres",
	Short: "Postgres database command", // Short description of the command
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()      // Show command help
		os.Exit(1)     // Exit with status code 1
	},
}

// gmGreenplumCmd represents the command for Greenplum database.
var gmGreenplumCmd = &cobra.Command{
	Use:   "greenplum",
	Short: "GreenplumCmd database command", // Short description of the command
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()      // Show command help
		os.Exit(1)     // Exit with status code 1
	},
}

// gmFakegresCmd represents the command for Fakegres database.
var gmFakegresCmd = &cobra.Command{
	Use:   "fakegres",
	Short: "Fakegres database command", // Short description of the command
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()      // Show command help
		os.Exit(1)     // Exit with status code 1
	},
}
