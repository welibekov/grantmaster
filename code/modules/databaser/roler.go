package databaser

import (
	"context"

	"github.com/welibekov/grantmaster/modules/role/types"
)

type Roler interface {
	Apply(context.Context, []types.Role) error
	Drop(context.Context, []types.Role) error
	Revoke(context.Context, []types.Role) error
	Grant(context.Context, []types.Role) error
	Get(context.Context) ([]types.Role, error)
}
