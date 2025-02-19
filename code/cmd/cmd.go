package cmd

import (
	"fmt"
	"os"

	"io/ioutil"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var configPath string

var rootCmd = &cobra.Command{
	Use:   "gm",
	Short: "Grantmaster CLI application",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Fprintln(os.Stderr, "No YAML structure provided")
			os.Exit(1)
		}

		if configPath != "" {
			config, err := readConfig(configPath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error reading config: %v\n", err)
				os.Exit(1)
			}

			fmt.Printf("Config: %+v\n", config)
		}

		yamlData := args[len(args)-1] // Last argument is the YAML structure.
		fmt.Printf("YAML Structure: %s\n", yamlData)
	},
}

func Execute() {
	rootCmd.PersistentFlags().StringVar(&configPath, "config", "", "Path to the config file")
	rootCmd.MarkPersistentFlagRequired("config")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// Helper Function to Read YAML Config
func readConfig(filePath string) (interface{}, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config interface{}
	err = yaml.Unmarshal(data, &config)
	return config, err
}
