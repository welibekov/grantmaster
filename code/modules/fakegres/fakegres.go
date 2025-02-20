package fakegres

import (
	"os"

	"github.com/welibekov/grantmaster/modules/types"
)

type Fakegres struct {
	rootDir string
}

// New creates a new instance of Fakegres with the provided configuration.
func New(config map[string]string) (*Fakegres, error) {
	fakegres := &Fakegres{}

	rootDir, found := config["GM_FAKEGRES_ROOTDIR"]
	if !found {
		rootDir = "/tmp/fakegres"
	}

	if _, err := os.Stat(rootDir); os.IsNotExist(err) {
		err := os.Mkdir(rootDir, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}

	fakegres.rootDir = rootDir

	return fakegres, nil
}

// Apply processes a slice of policies and applies the specified actions (grant/revoke).
func (f *Fakegres) Apply(policies []types.Policy) error {
	for _, policy := range policies {
		if err := f.apply(policy); err != nil {
			return err
		}
	}

	return nil
}
