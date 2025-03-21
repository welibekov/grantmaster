package policy

import "github.com/welibekov/grantmaster/modules/policy/base"

type FGPolicy struct {
	*base.Policy
}

func New(cfg map[string]string) (*FGPolicy, error) {
	basePolicy := base.NewPolicy(cfg)

	return &FGPolicy{
		Policy: basePolicy,
	}, nil
}
