package base

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type Runtest struct {
	Tests   []string
	ExecDir string

	gmBin string
}

func New(tests []string) (*Runtest, error) {
	gmBin, err := os.Executable()
	if err != nil {
		return nil, fmt.Errorf("couldn't get executable: %v", err)
	}

	return &Runtest{
		Tests:   tests,
		ExecDir: filepath.Dir(gmBin),

		gmBin: gmBin,
	}, nil
}

func (r *Runtest) notImplemented() error {
	return fmt.Errorf("NYI")
}

func (r *Runtest) Prepare() (func() error, error) {
	return func() error { return nil }, r.notImplemented()
}

func (r *Runtest) Execute() error {
	gmTestDir, err := os.MkdirTemp(os.TempDir(), "gm-runtest-")
	if err != nil {
		return fmt.Errorf("couldn't create test directory: %v", err)
	}

	env := []string{
		fmt.Sprintf("GM_BIN=%s", r.gmBin),
		fmt.Sprintf("GM_TEST_DIR=%s", gmTestDir),
	}

	var testErr error

	for _, test := range r.Tests {
		if err := r.exec(test, env).Run(); err != nil {
			testErr = err
		}
	}

	return testErr
}

func (r *Runtest) exec(test string, env []string) *exec.Cmd {
	cmd := exec.Command(test)
	cmd.Env = append(os.Environ(), env...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd
}
