package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/welibekov/grantmaster/internal/assets"
	"github.com/welibekov/grantmaster/internal/config"
	"github.com/welibekov/grantmaster/internal/policy"
	"github.com/welibekov/grantmaster/internal/policy/types"
	"github.com/welibekov/grantmaster/internal/policy/utils"
	"gopkg.in/yaml.v3"
)

// init initializes the commands for the application.
func init() {
	// Add the gmPolicyCmd to the specified commands.
	for _, gmCmd := range []*cobra.Command{gmPostgresCmd, gmFakegresCmd, gmGreenplumCmd} {
		gmCmd.AddCommand(gmPolicyCmd)
	}

	// Add specific policy commands to the gmPolicyCmd.
	for _, policyCmd := range []*cobra.Command{gmApplyPolicyCmd, gmGetPolicyCmd, gmEqualPolicyCmd} {
		gmPolicyCmd.AddCommand(policyCmd)
	}
}

// gmPolicyCmd is the command to manipulate database policies.
var gmPolicyCmd = &cobra.Command{
	Use:   "policy",
	Short: "Manipulate database policies",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help() // Show help if no subcommand is provided.
		os.Exit(1) // Exit with an error code.
	},
}

// gmApplyPolicyCmd is the command to apply policies from a specified YAML file or directory.
var gmApplyPolicyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply policies from the specified YAML file or directory",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Ensure that at least one argument (policy file or directory) is provided.
		if len(args) == 0 {
			return fmt.Errorf("policy.yaml file or directory not provided")
		}

		// Call the applyPolicy function with the provided argument.
		return applyPolicy(args[0])
	},
}

// applyPolicy applies policies from the given policy file.
func applyPolicy(policyFile string) error {
	// Load configuration from environment variables.
	cfg := config.Load()

	// Set context for database operations.
	ctx := context.Background()

	// Read policies from the specified file or directory.
	policies, err := assets.ReadAssets[types.Policy](policyFile)
	if err != nil {
		return fmt.Errorf("couldn't read policies: %v", err)
	}

	// Detect duplicated policies based on username.
	if err := assets.DetectDuplicated[types.Policy](policies, func(r types.Policy) string { return r.Username }); err != nil {
		return fmt.Errorf("duplicated policies found: %v", err)
	}

	// Validate role names against database role prefix criteria.
	if err := utils.CheckPrefix(policies, cfg[config.DatabaseRolePrefix]); err != nil {
		return fmt.Errorf("some role names do not satisfy %s criteria: %v", config.DatabaseRolePrefix, err)
	}

	// Create a new instance of the database with the provided configuration.
	databaseInstance, err := policy.New(ctx, cfg)
	if err != nil {
		return fmt.Errorf("failed to create database instance: %w", err)
	}

	// Apply the policies to the database.
	return databaseInstance.Apply(ctx, policies)
}

// gmGetPolicyCmd is the command to retrieve existing policies.
var gmGetPolicyCmd = &cobra.Command{
	Use:   "get",
	Short: "Get existing policies",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Set context for database operations.
		ctx := context.Background()

		// Create an instance of the database.
		databaseInstance, err := policy.New(ctx, config.Load())
		if err != nil {
			return fmt.Errorf("failed to create database instance: %w", err)
		}

		// Retrieve current policies from the database.
		policies, err := databaseInstance.Get(ctx)
		if err != nil {
			return fmt.Errorf("couldn't get existing policies: %v", err)
		}

		// Marshal the policies into YAML format.
		yamlBytes, err := yaml.Marshal(policies)
		if err != nil {
			return fmt.Errorf("couldn't marshal policies: %v", err)
		}

		// Write the YAML output to standard output.
		_, err = os.Stdout.Write(yamlBytes)

		return err // Return any error that occurred during writing.
	},
}

// gmEqualPolicyCmd is the command to compute the difference between two sets of roles.
var gmEqualPolicyCmd = &cobra.Command{
	Use:   "equal",                      // Command usage
	Short: "Find two policies equality", // Short description of command functionality
	RunE: func(cmd *cobra.Command, args []string) error {
		// Ensure exactly two arguments (file paths) are provided.
		if len(args) != 2 {
			return fmt.Errorf("please provide two role file paths") // Return an error if the count of arguments is incorrect
		}

		// Read the first set of policies from the specified file.
		policiesFirst, err := assets.ReadAssets[types.Policy](args[0])
		if err != nil {
			return fmt.Errorf("couldn't read policies from first file: %v", err) // Return error if reading the first policies fails
		}

		// Read the second set of policies from the specified file.
		policiesSecond, err := assets.ReadAssets[types.Policy](args[1])
		if err != nil {
			return fmt.Errorf("couldn't read policies from second file: %v", err) // Return error if reading the second policies fails
		}

		// Compare the two sets of policies for equality.
		if !utils.Equal(policiesFirst, policiesSecond) {
			os.Exit(1) // Exit with a non-zero status if the policies differ.
		}

		return nil // Return nil if the policies are equal.
	},
}
