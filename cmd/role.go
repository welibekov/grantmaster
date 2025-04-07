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

var roleFlags = struct {
	getAll bool // Flag to indicate if all roles should be retrieved.
}{}

// init sets up the command structure for managing roles by adding the role command to multiple commands.
func init() {
	// Iterate over predefined commands and add the role command to each.
	for _, gmCmd := range []*cobra.Command{gmPostgresCmd, gmFakegresCmd, gmGreenplumCmd} {
		gmCmd.AddCommand(gmRoleCmd) // Add the role command to each of the specified commands.
	}

	// Add specific role subcommands to the main role command.
	for _, roleCmd := range []*cobra.Command{gmApplyRoleCmd, gmGetRoleCmd, gmEqualRoleCmd} {
		gmRoleCmd.AddCommand(roleCmd) // Add subcommands to the gmRoleCmd.
	}

	// Define a persistent flag for retrieving all roles.
	gmGetRoleCmd.PersistentFlags().BoolVar(&roleFlags.getAll, "all", false, "Get all existing roles")
}

// gmRoleCmd defines the command for managing database roles.
var gmRoleCmd = &cobra.Command{
	Use:   "role",                      // Command usage.
	Short: "Manipulate database roles", // Short description of command functionality.
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help() // Display help message if no subcommand is provided.
		os.Exit(1) // Exit the application with an error status.
	},
}

// gmApplyRoleCmd defines the command to apply roles from a specified YAML file or directory.
var gmApplyRoleCmd = &cobra.Command{
	Use:   "apply",                                                 // Command usage.
	Short: "Apply roles from the specified YAML file or directory", // Short description of command functionality.
	RunE: func(cmd *cobra.Command, args []string) error {
		// Check if a role YAML file or directory is provided.
		if len(args) == 0 {
			return fmt.Errorf("role.yaml file or directory not provided") // Return an error if no arguments are supplied.
		}

		// Call applyRole function to apply the roles defined in the first argument.
		return applyRole(args[0])
	},
}

// applyRole applies roles defined in a specified file.
func applyRole(roleFile string) error {
	// Load configuration from environment variables.
	cfg := config.Load()

	// Create a background context for database operations.
	ctx := context.Background()

	// Read roles from the specified file or directory.
	roles, err := assets.ReadAssets[types.Role](roleFile)
	if err != nil {
		return fmt.Errorf("couldn't read roles: %v", err) // Return an error if reading roles fails.
	}

	// Detect duplicated roles in the loaded roles.
	if err := assets.DetectDuplicated[types.Role](roles, func(r types.Role) string { return r.Name }); err != nil {
		return fmt.Errorf("duplicated roles found: %v", err) // Return an error if duplicates are found.
	}

	// Validate role names against the specified prefix criteria from the configuration.
	if err := utils.CheckPrefix(roles, cfg[config.DatabaseRolePrefix]); err != nil {
		return fmt.Errorf("some role names do not satisfy %s criteria: %v", config.DatabaseRolePrefix, err) // Return error if prefix check fails.
	}

	// Create an instance of the database object.
	databaseInstance, err := role.New(ctx, cfg)
	if err != nil {
		return fmt.Errorf("failed to create database instance: %w", err) // Return error if database instance creation fails.
	}

	// Apply the roles to the database.
	return databaseInstance.Apply(ctx, roles) // Return any error occurred during the application of roles.
}

// gmGetRoleCmd defines the command to get existing roles from the database.
var gmGetRoleCmd = &cobra.Command{
	Use:   "get",                // Command usage.
	Short: "Get existing roles", // Short description of command functionality.
	RunE: func(cmd *cobra.Command, args []string) error {
		// Create a background context for database operations.
		ctx := context.Background()

		// Load configuration from environment variables.
		cfg := config.Load()

		// Adjust the role prefix in the configuration if the "get all" flag is set.
		if roleFlags.getAll {
			cfg[config.DatabaseRolePrefix] = "_" // by SQL definition "_" applies all values.
		}

		// Create an instance of the database object.
		databaseInstance, err := role.New(ctx, cfg)
		if err != nil {
			return fmt.Errorf("failed to create database instance: %w", err) // Return error if database instance creation fails.
		}

		// Retrieve roles from the database.
		roles, err := databaseInstance.Get(ctx)
		if err != nil {
			return fmt.Errorf("couldn't get existing roles: %v", err) // Return error if getting roles fails.
		}

		// Marshal the retrieved roles into YAML format.
		yamlBytes, err := yaml.Marshal(roles)
		if err != nil {
			return fmt.Errorf("couldn't marshal roles: %v", err) // Return error if YAML marshaling fails.
		}

		// Write the YAML output to standard output.
		_, err = os.Stdout.Write(yamlBytes)

		return err // Return any error occurred during writing.
	},
}

// gmEqualRoleCmd defines the command to compute the difference between two sets of roles.
var gmEqualRoleCmd = &cobra.Command{
	Use:   "equal",                    // Command usage.
	Short: "Find two roles' equality", // Short description of command functionality.
	RunE: func(cmd *cobra.Command, args []string) error {
		// Ensure exactly two arguments are provided.
		if len(args) != 2 {
			return fmt.Errorf("please provide two role file paths") // Return an error if the count of arguments is incorrect.
		}

		// Read the first set of roles from the specified file or directory.
		rolesFirst, err := assets.ReadAssets[types.Role](args[0])
		if err != nil {
			return fmt.Errorf("couldn't read roles from first file: %v", err) // Return error if reading the first roles fails.
		}

		// Read the second set of roles from the specified file or directory.
		rolesSecond, err := assets.ReadAssets[types.Role](args[1])
		if err != nil {
			return fmt.Errorf("couldn't read roles from second file: %v", err) // Return error if reading the second roles fails.
		}

		// Compare the two sets of roles for equality.
		if !utils.Equal(rolesFirst, rolesSecond) {
			os.Exit(1) // Exit with a non-zero status if the roles differ.
		}

		return nil // Return nil if roles are equal.
	},
}
