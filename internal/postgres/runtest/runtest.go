package runtest

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/welibekov/grantmaster/internal/runtest/base"
	"github.com/welibekov/grantmaster/internal/types"
)

// Runtest represents a test runner that extends the base Runtest functionality.
type Runtest struct {
	*base.Runtest // Embed base Runtest to inherit its methods and properties.
}

// New initializes a new Runtest instance with the provided test names.
// It sets up a base Runtest with the Postgres type and the given tests.
func New(tests []string) (*Runtest, error) {
	rt := &Runtest{}

	// Create a new base Runtest and handle potential errors.
	baseRuntest, err := base.New(types.Postgres, tests)
	if err != nil {
		return nil, fmt.Errorf("couldn't prepare base runtest: %v", err)
	}

	// Assign the base Runtest to the Runtest instance.
	rt.Runtest = baseRuntest

	return rt, nil
}

// Prepare sets up the environment for running tests by executing preparation scripts.
// It returns cleanup function and error if any occurred during setup.
func (r *Runtest) Prepare() (func() error, error) {
	// Construct the path to the preparation scripts.
	run := filepath.Join(filepath.Dir(r.ExecDir), "tests/postgres/prepare.sh")
	spinUp := run + " spinup"    // Command to spin up the test environment.
	spinDown := run + " spindown" // Command to spin down the test environment.

	// Return a cleanup function and execute the spin-up command.
	return func() error { return r.exec(spinDown) }, r.exec(spinUp)
}

// exec runs a shell command using bash and returns any error encountered.
// It captures standard input, output, and error streams to/from the console.
func (r *Runtest) exec(command string) error {
	cmd := exec.Command("bash", "-c", command)
	cmd.Stderr = os.Stderr // Direct stderr output to the console.
	cmd.Stdout = os.Stdout // Direct stdout output to the console.
	cmd.Stdin = os.Stdin   // Allow input to come from the console.

	return cmd.Run() // Execute the command and return any error.
}
