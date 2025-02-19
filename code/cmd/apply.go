package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/welibekov/grantmaster/modules/database"
	"github.com/welibekov/grantmaster/modules/types"
	"gopkg.in/yaml.v3"
)

var policyFile string

func init() {
	rootCmd.AddCommand(gmCmd)
	gmCmd.Flags().StringVar(&policyFile, "policy", "", "Path to policy YAML file (mandatory)")
	gmCmd.MarkFlagRequired("policy")
}

var gmCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply policies from the specified YAML file",
	Run: func(cmd *cobra.Command, args []string) {
		if err := run(); err != nil {
			fmt.Println("Error:", err)
		}
	},
}

func run() error {
	// Load configuration from environment variables
	config := loadConfig()

	// Read policy YAML file
	data, err := ioutil.ReadFile(policyFile)
	if err != nil {
		return fmt.Errorf("failed to read policy file: %w", err)
	}

	var policies []types.Policy
	if err := yaml.Unmarshal(data, &policies); err != nil {
		return fmt.Errorf("failed to unmarshal YAML: %w", err)
	}

	// Create an instance of database
	databaseInstance, err := database.New(config)
	if err != nil {
		return fmt.Errorf("failed to create database instance: %w", err)
	}

	// Apply policies
	return databaseInstance.Apply(policies)
}

func loadConfig() map[string]string {
	config := make(map[string]string)
	for _, env := range os.Environ() {
		kv := strings.SplitN(env, "=", 2)
		if len(kv) == 2 && strings.HasPrefix(kv[0], "GM_") {
			config[kv[0]] = kv[1]
		}
	}
	return config
}
