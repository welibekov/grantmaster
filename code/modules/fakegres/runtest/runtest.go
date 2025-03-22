package runtest

import (
	"fmt"

	"github.com/welibekov/grantmaster/modules/runtest/base"
	"github.com/welibekov/grantmaster/modules/types"
)

type Runtest struct {
	*base.Runtest
}

func New(tests []string) (*Runtest, error) {
	rt := &Runtest{}

	baseRuntest, err := base.New(types.Postgres, tests)
	if err != nil {
		return nil, fmt.Errorf("couldn't prepare base runtest: %v", err)
	}

	rt.Runtest = baseRuntest

	return rt, nil
}
