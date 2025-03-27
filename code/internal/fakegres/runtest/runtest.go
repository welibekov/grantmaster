package runtest

import (
	"fmt"

	"github.com/welibekov/grantmaster/internal/runtest/base"
	"github.com/welibekov/grantmaster/internal/types"
)

// Runtest represents a test runner that embeds the base Runtest functionality.
type Runtest struct {
	*base.Runtest
}

// New creates a new instance of Runtest with the specified test names.
// It initializes the embedded base Runtest and returns an error if the setup fails.
func New(tests []string) (*Runtest, error) {
	rt := &Runtest{} // Initialize a new Runtest instance.

	// Create a new base Runtest using the specified database type and test names.
	baseRuntest, err := base.New(types.Postgres, tests)
	if err != nil {
		// Return an error if the base Runtest cannot be prepared.
		return nil, fmt.Errorf("couldn't prepare base runtest: %v", err)
	}

	rt.Runtest = baseRuntest // Assign the initialized base Runtest to the Runtest instance.

	return rt, nil // Return the new Runtest instance.
}
