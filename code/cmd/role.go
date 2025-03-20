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
	"github.com/welibekov/grantmaster/modules/role/utils"

	"gopkg.in/yaml.v3"
)

func init() {
	rootCmd.AddCommand(gmRoleCmd)
	gmRoleCmd.AddCommand(gmApplyRoleCmd)
	gmRoleCmd.AddCommand(gmGetRoleCmd)
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
	cfg := config.Load()

	// Create an instance of database
	ctx := context.Background()

	// Read roles from file or directory.
	roles, err := assets.ReadAssets[types.Role](roleFile)
	if err != nil {
		return fmt.Errorf("couldn't read roles: %v", err)
	}

	// Detect duplicated roles
	if err := assets.DetectDuplicated[types.Role](roles, func(r types.Role) string { return r.Name }); err != nil {
		return fmt.Errorf("duplicated roles found: %v", err)
	}

	// Detect roles that doesn't meat prefix criteria.
	if err := utils.CheckPrefix(roles, cfg[config.DatabaseRolePrefix]); err != nil {
		return fmt.Errorf("some role names are not satisfy GM_ROLE_PREFIX criteria: %v", err)
	}

	databaseInstance, err := role.New(ctx, cfg)
	if err != nil {
		return fmt.Errorf("failed to create database instance: %w", err)
	}

	// Apply roles
	return databaseInstance.Apply(ctx, roles)
}

// Get policies
var gmGetRoleCmd = &cobra.Command{
	Use:   "get",
	Short: "Get existing roles",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Create an instance of database
		ctx := context.Background()

		databaseInstance, err := role.New(ctx, config.Load())
		if err != nil {
			return fmt.Errorf("failed to create database instance: %w", err)
		}

		// Get roles
		roles, err := databaseInstance.Get(ctx)
		if err != nil {
			return fmt.Errorf("couldn't get existing roles: %v", err)
		}

		yamlBytes, err := yaml.Marshal(roles)
		if err != nil {
			return fmt.Errorf("couldn't marshal roles: %v", err)
		}

		fmt.Println(string(yamlBytes))

		return nil
	},
}
