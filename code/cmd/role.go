package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/welibekov/grantmaster/modules/assets"
	"github.com/welibekov/grantmaster/modules/config"
	"github.com/welibekov/grantmaster/modules/role"
	"github.com/welibekov/grantmaster/modules/role/types"
)

func init() {
	rootCmd.AddCommand(gmRoleCmd)
	gmRoleCmd.AddCommand(gmApplyRoleCmd)
}

var gmRoleCmd = &cobra.Command{
	Use:   "role",
	Short: "Maniulate database roles",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(1)
	},
}

var gmApplyRoleCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply roles from the specified YAML file or directory",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("role.yaml file or directory not provided")
		}

		return applyRole(args[0])
	},
}

func applyRole(roleFile string) error {
	// Load configuration from environment variables
	config := config.Load()

	// Read roles from file or directory.
	roles, err := assets.ReadAssets[types.Role](roleFile)
	if err != nil {
		return fmt.Errorf("couldn't read roles: %v", err)
	}

	// Detect duplicated roles
	if err := assets.DetectDuplicated[types.Role](roles, func(r types.Role) string { return r.Name }); err != nil {
		return fmt.Errorf("duplicated roles found: %v", err)
	}

	// Create an instance of database
	ctx := context.Background()

	databaseInstance, err := role.New(ctx, config)
	if err != nil {
		return fmt.Errorf("failed to create database instance: %w", err)
	}

	// Apply roles
	return databaseInstance.Apply(ctx, roles)
}
