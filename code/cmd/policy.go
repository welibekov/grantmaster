package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/welibekov/grantmaster/modules/assets"
	"github.com/welibekov/grantmaster/modules/config"
	"github.com/welibekov/grantmaster/modules/database"
	"github.com/welibekov/grantmaster/modules/policy/types"
)

var policyFile string

func init() {
	gmApplyCmd.AddCommand(gmApplyPolicyCmd)
	gmApplyPolicyCmd.Flags().StringVar(&policyFile, "policy", "", "Path to policy YAML file (mandatory)")
	gmApplyPolicyCmd.MarkFlagRequired("policy")
}

var gmApplyPolicyCmd = &cobra.Command{
	Use:   "policy",
	Short: "Apply policies from the specified YAML file or directory",
	Run: func(cmd *cobra.Command, args []string) {
		if err := applyPolicy(); err != nil {
			fmt.Println("Error:", err)
		}
	},
}

func applyPolicy() error {
	// Load configuration from environment variables
	config := config.Load()

	// Read policies from file or directory.
	policies, err := assets.ReadAssets[types.Policy](policyFile)
	if err != nil {
		return fmt.Errorf("couldn't read policies: %v", err)
	}

	// Detect duplicated policies
	if err := assets.DetectDuplicated[types.Policy](policies, func(r types.Policy) string { return r.Username }); err != nil {
		return fmt.Errorf("duplicated policies found: %v", err)
	}

	// Create an instance of database
	databaseInstance, err := database.New(config)
	if err != nil {
		return fmt.Errorf("failed to create database instance: %w", err)
	}

	// Apply policies
	return databaseInstance.ApplyPolicy(context.Background(), policies)
}
