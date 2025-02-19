package fakegres

import (
	"fmt"

	"github.com/welibekov/grantmaster/modules/types"
)

type applyFunc func(string, []string) error

type Fakegres struct{}

// New creates a new instance of Fakegres with the provided configuration.
func New(config map[string]string) (*Fakegres, error) {
	return &Fakegres{}, nil
}

// Apply processes a slice of policies and applies the specified actions (grant/revoke).
func (f *Fakegres) Apply(policies []types.Policy) error {
	for _, policy := range policies {
		for _, action := range policy.Actions {
			var apply applyFunc

			switch action.Action {
			case "grant":
				apply = f.Grant
			case "revoke":
				apply = f.Revoke
			}

			if err := apply(policy.Username, action.Roles); err != nil {
				return err
			}
		}
	}

	return fmt.Errorf("NYI")
}

// Grant applies a grant action for the given username and roles.
func (f *Fakegres) Grant(string, []string) error {
	return fmt.Errorf("NYI")
}

// Revoke applies a revoke action for the given username and roles.
func (f *Fakegres) Revoke(string, []string) error {
	return fmt.Errorf("NYI")
}
