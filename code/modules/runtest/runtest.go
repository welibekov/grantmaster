package runtest

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/welibekov/grantmaster/modules/types"
)

type Runtest struct {
	Type  types.DatabaseType
	Tests []string
}

func New(config map[string]string, tests ...string) *Runtest {
	databaseType := types.DatabaseType(config["GM_DATABASE_TYPE"])

	if databaseType == "" {
		databaseType = types.Fakegres
	}

	logrus.Debugln("database type:", databaseType.ToString())

	return &Runtest{
		Tests: tests,
		Type:  databaseType,
	}
}

func (r *Runtest) Execute() error {
	gmBin, err := os.Executable()
	if err != nil {
		return fmt.Errorf("couldn't get executable: %v", err)
	}

	env := []string{
		fmt.Sprintf("GM_BIN=%s", gmBin),
	}

	cancel, err := r.prepareTestEnv(filepath.Dir(gmBin))
	if err != nil {
		return fmt.Errorf("couldn't prepare test env for %s: %v", r.Type.ToString(), err)
	}

	defer cancel()

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
