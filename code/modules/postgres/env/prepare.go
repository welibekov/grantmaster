package env

import (
	"os"
	"os/exec"
	"path/filepath"
)

func Prepare(execDir string) (func() error, error) {
	run := filepath.Join(filepath.Dir(execDir), "tests/postgres/prepare.sh")
	spinUp := run + " spinup"
	spinDown := run + " spindown"

	return func() error { return execCommand(spinDown) }, execCommand(spinUp)
}

func execCommand(command string) error {
	cmd := exec.Command("bash", "-c", command)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	return cmd.Run()
}
