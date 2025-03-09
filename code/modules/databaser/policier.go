package databaser

import (
	"context"

	"github.com/welibekov/grantmaster/modules/policy/types"
)

type Policier interface {
	Apply(context.Context, []types.Policy) error
	Grant(context.Context, []types.Policy) error
	Revoke(context.Context, []types.Policy) error
	Get(context.Context) ([]types.Policy, error)
}
