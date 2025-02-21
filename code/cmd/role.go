package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	gmApplyCmd.AddCommand(gmApplyRoleCmd)
	gmApplyRoleCmd.Flags().StringVar(&policyFile, "policy", "", "Path to role YAML file or directory (mandatory)")
	gmApplyRoleCmd.MarkFlagRequired("role")
}

var gmApplyRoleCmd = &cobra.Command{
	Use:   "role",
	Short: "Apply roles from the specified YAML file or directory",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("NYI")
		os.Exit(1)
	},
}
