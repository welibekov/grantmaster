package base

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/sirupsen/logrus"
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
	env := []string{
		fmt.Sprintf("GM_BIN=%s", r.gmBin),
	}

	for _, test := range r.Tests {
		if err := r.exec(test, env).Run(); err != nil {
			logrus.Warnf("test '%s' failed: %v", test, err)
		}
	}

	return nil
}

func (r *Runtest) exec(test string, env []string) *exec.Cmd {
	cmd := exec.Command(test)
	cmd.Env = append(os.Environ(), env...)
	cmd.Stdout = os.Stdout

	return cmd
}
