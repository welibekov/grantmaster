package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/welibekov/grantmaster/modules/assets"
	"github.com/welibekov/grantmaster/modules/config"
	"github.com/welibekov/grantmaster/modules/role"
	"github.com/welibekov/grantmaster/modules/role/types"
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
