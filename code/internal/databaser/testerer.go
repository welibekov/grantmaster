package databaser

// RunTesterer defines the methods required for a test runner.
type RunTesterer interface {
    // Prepare prepares the test environment and returns a function
    // to clean up afterward, along with any error that occurred.
    Prepare() (func() error, error)

    // Execute runs the test and returns any error that occurred during execution.
    Execute() error
}
