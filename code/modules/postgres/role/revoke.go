package role

import (
	"context"

	"github.com/welibekov/grantmaster/modules/role/types"
)

func (p *PGRole) Revoke(ctx context.Context, roles []types.Role) error {
	return p.Drop(ctx, roles)
}
