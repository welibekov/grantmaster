package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/welibekov/grantmaster/internal/assets"
	"github.com/welibekov/grantmaster/internal/config"
	"github.com/welibekov/grantmaster/internal/role"
	"github.com/welibekov/grantmaster/internal/role/types"
	"github.com/welibekov/grantmaster/internal/role/utils"

	"gopkg.in/yaml.v3"
)

// Initialize the command structure for roles
func init() {
	for _, gmCmd := range []*cobra.Command{gmPostgresCmd, gmFakegresCmd, gmGreenplumCmd} {
		gmCmd.AddCommand(gmRoleCmd)
	}

	for _, roleCmd := range []*cobra.Command{gmApplyRoleCmd, gmGetRoleCmd, gmEqualRoleCmd} {
		gmRoleCmd.AddCommand(roleCmd)
	}
}

// Define the command for managing roles
var gmRoleCmd = &cobra.Command{
	Use:   "role",                      // Command usage
	Short: "Manipulate database roles", // Short description of command functionality
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help() // Display help message if no subcommand is provided
		os.Exit(1) // Exit the application with an error status
	},
}

// Define the command to apply roles
var gmApplyRoleCmd = &cobra.Command{
	Use:   "apply",                                                 // Command usage
	Short: "Apply roles from the specified YAML file or directory", // Short description of command functionality
	RunE: func(cmd *cobra.Command, args []string) error {
		// Check if a role YAML file or directory is provided
		if len(args) == 0 {
			return fmt.Errorf("role.yaml file or directory not provided") // Return an error if no arguments are supplied
		}

		// Call applyRole function to apply the roles defined in the first argument
		return applyRole(args[0])
	},
}

// applyRole applies roles defined in a specified file
func applyRole(roleFile string) error {
	// Load configuration from environment variables
	cfg := config.Load()

	// Create a background context for database operations
	ctx := context.Background()

	// Read roles from the specified file or directory
	roles, err := assets.ReadAssets[types.Role](roleFile)
	if err != nil {
		return fmt.Errorf("couldn't read roles: %v", err) // Return an error if reading roles fails
	}

	// Detect duplicated roles in the loaded roles
	if err := assets.DetectDuplicated[types.Role](roles, func(r types.Role) string { return r.Name }); err != nil {
		return fmt.Errorf("duplicated roles found: %v", err) // Return an error if duplicates are found
	}

	// Validate role names against the specified prefix criteria from the configuration
	if err := utils.CheckPrefix(roles, cfg[config.DatabaseRolePrefix]); err != nil {
		return fmt.Errorf("some role names are not satisfy GM_ROLE_PREFIX criteria: %v", err) // Return error if prefix check fails
	}

	// Create an instance of the database object
	databaseInstance, err := role.New(ctx, cfg)
	if err != nil {
		return fmt.Errorf("failed to create database instance: %w", err) // Return error if database instance creation fails
	}

	// Apply the roles to the database
	return databaseInstance.Apply(ctx, roles)
}

// Define the command to get existing roles
var gmGetRoleCmd = &cobra.Command{
	Use:   "get",                // Command usage
	Short: "Get existing roles", // Short description of command functionality
	RunE: func(cmd *cobra.Command, args []string) error {
		// Create a background context for database operations
		ctx := context.Background()

		// Create an instance of the database object
		databaseInstance, err := role.New(ctx, config.Load())
		if err != nil {
			return fmt.Errorf("failed to create database instance: %w", err) // Return error if database instance creation fails
		}

		// Retrieve roles from the database
		roles, err := databaseInstance.Get(ctx)
		if err != nil {
			return fmt.Errorf("couldn't get existing roles: %v", err) // Return error if getting roles fails
		}

		// Marshal the retrieved roles into YAML format
		yamlBytes, err := yaml.Marshal(roles)
		if err != nil {
			return fmt.Errorf("couldn't marshal roles: %v", err) // Return error if YAML marshaling fails
		}

		// Write the YAML output to standard output
		_, err = os.Stdout.Write(yamlBytes)

		return err // Return any error occurred during writing
	},
}

// Define the command to compute the difference between two sets of roles
var gmEqualRoleCmd = &cobra.Command{
	Use:   "equal",                   // Command usage
	Short: "Find two roles equality", // Short description of command functionality
	RunE: func(cmd *cobra.Command, args []string) error {
		// Ensure exactly two arguments are provided
		if len(args) != 2 {
			return fmt.Errorf("please provide two role file paths") // Return an error if the count of arguments is incorrect
		}

		// Read the first set of roles from the specified file or directory
		rolesFirst, err := assets.ReadAssets[types.Role](args[0])
		if err != nil {
			return fmt.Errorf("couldn't read roles from first file: %v", err) // Return error if reading the first roles fails
		}

		// Read the second set of roles from the specified file or directory
		rolesSecond, err := assets.ReadAssets[types.Role](args[1])
		if err != nil {
			return fmt.Errorf("couldn't read roles from second file: %v", err) // Return error if reading the second roles fails
		}

		// Compare the two sets of roles for equality
		if !utils.Equal(rolesFirst, rolesSecond) {
			os.Exit(1) // Exit with a non-zero status if the roles differ
		}

		return nil // Return nil if roles are equal
	},
}
