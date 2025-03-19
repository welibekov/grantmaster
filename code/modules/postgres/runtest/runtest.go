package runtest

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/welibekov/grantmaster/modules/runtest/base"
)

type Runtest struct {
	*base.Runtest
}

func New(tests []string) (*Runtest, error) {
	rt := &Runtest{}

	baseRuntest, err := base.New(tests)
	if err != nil {
		return nil, fmt.Errorf("couldn't prepare base runtest: %v", err)
	}

	rt.Runtest = baseRuntest

	return rt, nil
}

func (r *Runtest) Prepare() (func() error, error) {
	run := filepath.Join(filepath.Dir(r.ExecDir), "tests/postgres/prepare.sh")
	spinUp := run + " spinup"
	spinDown := run + " spindown"

	return func() error { return r.exec(spinDown) }, r.exec(spinUp)
}

func (r *Runtest) exec(command string) error {
	cmd := exec.Command("bash", "-c", command)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	return cmd.Run()
}
