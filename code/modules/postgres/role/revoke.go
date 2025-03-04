package role

import (
	"context"

	"github.com/welibekov/grantmaster/modules/role/types"
)

func (p *PGRole) Revoke(ctx context.Context, roles []types.Role) error {
	return p.exec(ctx, roles, "postgres/role/revoke.tmpl")
}
