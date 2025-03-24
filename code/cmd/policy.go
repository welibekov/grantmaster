package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/welibekov/grantmaster/modules/assets"
	"github.com/welibekov/grantmaster/modules/config"
	"github.com/welibekov/grantmaster/modules/policy"
	"github.com/welibekov/grantmaster/modules/policy/types"
	"github.com/welibekov/grantmaster/modules/policy/utils"
	"gopkg.in/yaml.v3"
)

func init() {
	for _, gmCmd := range []*cobra.Command{gmPostgresCmd, gmFakegresCmd, gmGreenplumCmd} {
		gmCmd.AddCommand(gmPolicyCmd)
	}

	for _, policyCmd := range []*cobra.Command{gmApplyPolicyCmd, gmGetPolicyCmd, gmEqualPolicyCmd} {
		gmPolicyCmd.AddCommand(policyCmd)
	}
}

var gmPolicyCmd = &cobra.Command{
	Use:   "policy",
	Short: "Maniulate database policies",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(1)
	},
}

var gmApplyPolicyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply policies from the specified YAML file or directory",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("role.yaml file or directory not provided")
		}

		return applyPolicy(args[0])
	},
}

func applyPolicy(policyFile string) error {
	// Load configuration from environment variables
	cfg := config.Load()

	// Set context
	ctx := context.Background()

	// Read policies from file or directory.
	policies, err := assets.ReadAssets[types.Policy](policyFile)
	if err != nil {
		return fmt.Errorf("couldn't read policies: %v", err)
	}

	// Detect duplicated policies
	if err := assets.DetectDuplicated[types.Policy](policies, func(r types.Policy) string { return r.Username }); err != nil {
		return fmt.Errorf("duplicated policies found: %v", err)
	}

	// Detect roles that doesn't meat prefix criteria.
	if err := utils.CheckPrefix(policies, cfg[config.DatabaseRolePrefix]); err != nil {
		return fmt.Errorf("some role names are not satisfy GM_ROLE_PREFIX criteria: %v", err)
	}

	// Create an instance of database
	databaseInstance, err := policy.New(ctx, cfg)
	if err != nil {
		return fmt.Errorf("failed to create database instance: %w", err)
	}

	// Apply policies
	return databaseInstance.Apply(ctx, policies)
}

// Get policies
var gmGetPolicyCmd = &cobra.Command{
	Use:   "get",
	Short: "Get existing policies",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Create an instance of database
		ctx := context.Background()

		databaseInstance, err := policy.New(ctx, config.Load())
		if err != nil {
			return fmt.Errorf("failed to create database instance: %w", err)
		}

		// Get policies
		policies, err := databaseInstance.Get(ctx)
		if err != nil {
			return fmt.Errorf("couldn't get existing policies: %v", err)
		}

		yamlBytes, err := yaml.Marshal(policies)
		if err != nil {
			return fmt.Errorf("couldn't marshal policies: %v", err)
		}

		_, err = os.Stdout.Write(yamlBytes)

		return err
	},
}

// Define the command to compute the difference between two sets of roles
var gmEqualPolicyCmd = &cobra.Command{
	Use:   "equal",                      // Command usage
	Short: "Find two policies equality", // Short description of command functionality
	RunE: func(cmd *cobra.Command, args []string) error {
		// Ensure exactly two arguments are provided
		if len(args) != 2 {
			return fmt.Errorf("please provide two role file paths") // Return an error if the count of arguments is incorrect
		}

		// Read the first set of policies from the specified file or directory
		policiesFirst, err := assets.ReadAssets[types.Policy](args[0])
		if err != nil {
			return fmt.Errorf("couldn't read policies from first file: %v", err) // Return error if reading the first policies fails
		}

		// Read the second set of policies from the specified file or directory
		policiesSecond, err := assets.ReadAssets[types.Policy](args[1])
		if err != nil {
			return fmt.Errorf("couldn't read policies from second file: %v", err) // Return error if reading the second policies fails
		}

		// Compare the two sets of policies for equality
		if !utils.Equal(policiesFirst, policiesSecond) {
			os.Exit(1) // Exit with a non-zero status if the policies differ
		}

		return nil // Return nil if policies are equal
	},
}
