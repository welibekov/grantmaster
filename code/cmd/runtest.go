package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/welibekov/grantmaster/modules/config"
	"github.com/welibekov/grantmaster/modules/runtest"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(gmRuntestCmd)
}

var gmRuntestCmd = &cobra.Command{
	Use:   "runtest",
	Short: "Run tests of specific test type.",
	RunE: func(cmd *cobra.Command, args []string) error {
		tests := []string{}

		for _, arg := range args {
			err := filepath.WalkDir(arg, func(path string, d os.DirEntry, err error) error {
				if err != nil {
					return err
				}

				if !d.IsDir() && strings.HasSuffix(d.Name(), ".bash") {
					tests = append(tests, path)
				}

				return nil
			})

			if err != nil {
				fmt.Println("Error:", err)
			}
		}

		if len(tests) == 0 {
			return fmt.Errorf("no tests files found")
		}

		rt, err := runtest.New(config.Load(), tests)
		if err != nil {
			return err
		}

		cancel, err := rt.Prepare()
		if err != nil {
			return fmt.Errorf("couldn't prepare runtest env: %v", err)
		}

		defer cancel()

		return rt.Execute()
	},
}
