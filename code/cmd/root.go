package cmd

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.PersistentFlags().String("loglevel", "info", "Set loglevel")
}

var rootCmd = &cobra.Command{
	Use:   "gm",
	Short: "Grant Master CLI tool",
	Long:  `A command line tool to apply database permissions based on policies.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		level, err := cmd.Flags().GetString("loglevel")
		if err != nil {
			return err
		}

		return setLogLevel(level)
	},
}

func setLogLevel(level string) error {
	parsed, err := logrus.ParseLevel(level)
	if err != nil {
		return fmt.Errorf("invalid loglevel %s: %v", level, err)
	}

	logrus.SetLevel(parsed)

	return nil
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
