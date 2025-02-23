package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var roleFile string

func init() {
	gmApplyCmd.AddCommand(gmApplyRoleCmd)
	gmApplyRoleCmd.Flags().StringVar(&roleFile, "role", "", "Path to role YAML file or directory (mandatory)")
	gmApplyRoleCmd.MarkFlagRequired("role")
}

var gmApplyRoleCmd = &cobra.Command{
	Use:   "role",
	Short: "Apply roles from the specified YAML file or directory",
	Run: func(cmd *cobra.Command, args []string) {
		if err := applyRole(); err != nil {
			fmt.Println("Error:", err)
		}
	},
}

func applyRole() error {
	// Load configuration from environment variables
	// config := config.Load()

	return fmt.Errorf("NYI")
}
