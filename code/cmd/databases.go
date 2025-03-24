package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func init() {
	for _, gmCmd := range []*cobra.Command{
		gmPostgresCmd, gmFakegresCmd, gmGreenplumCmd,
	} {
		rootCmd.AddCommand(gmCmd)
	}
}

var gmPostgresCmd = &cobra.Command{
	Use:   "postgres",
	Short: "Postgres database command",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(1)
	},
}

var gmGreenplumCmd = &cobra.Command{
	Use:   "greenplum",
	Short: "GreenplumCmd database command",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(1)
	},
}

var gmFakegresCmd = &cobra.Command{
	Use:   "fakegres",
	Short: "Fakegres database command",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(1)
	},
}
