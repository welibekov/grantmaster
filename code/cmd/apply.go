package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var policyFile string

func init() {
	rootCmd.AddCommand(gmApplyCmd)
}

var gmApplyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply policies,roles from the specified YAML file or directory",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(1)
	},
}
