package role

import (
	"context"

	"github.com/welibekov/grantmaster/modules/role/types"
)

func (p *PGRole) Drop(ctx context.Context, roles []types.Role) error {
	return p.exec(ctx, roles, "postgres/role/drop.tmpl")
}
