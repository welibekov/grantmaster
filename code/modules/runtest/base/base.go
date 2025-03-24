package base

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/welibekov/grantmaster/modules/types"
)

// Runtest is responsible for preparing and executing test cases.
type Runtest struct {
	DatabaseType types.DatabaseType // Type of database being tested
	Tests        []string           // List of test cases to execute
	ExecDir      string             // Directory where the executable is located

	gmBin string // Path to the grantmaster binary
}

// New creates a new Runtest instance, returns an error if the executable path cannot be determined.
func New(dbType types.DatabaseType, tests []string) (*Runtest, error) {
	gmBin, err := os.Executable()
	if err != nil {
		return nil, fmt.Errorf("couldn't get executable: %v", err)
	}

	return &Runtest{
		DatabaseType: dbType,
		Tests:        tests,
		ExecDir:      filepath.Dir(gmBin),

		gmBin: gmBin,
	}, nil
}

// notImplemented returns an error indicating that the method is not yet implemented.
func (r *Runtest) notImplemented() error {
	return fmt.Errorf("NYI") // 'NYI' stands for 'Not Yet Implemented'
}

// Prepare prepares the environment for executing the tests. 
// This function is expected to return a cleanup function and an error if applicable.
// Currently, it does not implement any preparation logic, hence NYI error.
func (r *Runtest) Prepare() (func() error, error) {
	return func() error { return nil }, r.notImplemented() // No preparation logic implemented yet
}

// Execute runs the tests defined in the Runtest instance.
// It creates a temporary directory for test execution and uses the exec.Command to run each test.
func (r *Runtest) Execute() error {
	gmTestDir, err := os.MkdirTemp(os.TempDir(), fmt.Sprintf("gm-runtest-%s-", r.DatabaseType.ToString()))
	if err != nil {
		return fmt.Errorf("couldn't create test directory: %v", err)
	}

	// Setting up environment variables for the test execution
	env := []string{
		fmt.Sprintf("GM_BIN=%s", r.gmBin),           // Path to the grantmaster binary
		fmt.Sprintf("GM_TEST_DIR=%s", gmTestDir),     // Path to the test directory
	}

	var testErr error // Variable to track the last encountered test error

	// Iterating through the list of tests to execute them
	for _, test := range r.Tests {
		if err := r.exec(test, env).Run(); err != nil {
			testErr = err // Updating testErr with the latest encountered error
		}
	}

	return testErr // Returning the last error encountered, if any
}

// exec constructs and returns an exec.Cmd for the given test case.
func (r *Runtest) exec(test string, env []string) *exec.Cmd {
	cmd := exec.Command(test)           // Create a new command for the specified test
	cmd.Env = append(os.Environ(), env...) // Append the provided environment variables to the command
	cmd.Stdin = os.Stdin               // Set standard input for the command
	cmd.Stdout = os.Stdout             // Set standard output for the command
	cmd.Stderr = os.Stderr             // Set standard error for the command

	return cmd // Returning the constructed command
}
