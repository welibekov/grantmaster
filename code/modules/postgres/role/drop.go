package role

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/welibekov/grantmaster/modules/role/types"
)

func (p *PGRole) Drop(ctx context.Context, roles []types.Role) error {
	var query string

	for _, role := range roles {
		for _, schema := range role.Schemas {
			query += fmt.Sprintf("REVOKE ALL PRIVILEGES ON SCHEMA %s FROM %s;", schema.Schema, role.Name)
			query += fmt.Sprintf("REVOKE GRANT OPTION FOR ALL PRIVILEGES ON SCHEMA %s FROM %s;", schema.Schema, role.Name)
			query += fmt.Sprintf("REVOKE ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA %s FROM %s;", schema.Schema, role.Name)
			query += fmt.Sprintf("REVOKE ALL PRIVILEGES ON ALL FUNCTIONS IN SCHEMA %s FROM %s;", schema.Schema, role.Name)
			query += fmt.Sprintf("DROP ROLE %s;", role.Name)
		}
	}

	logrus.Debugln(query) // Log the generated query for debugging purposes

	_, err := p.pool.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("couldn't execute drop query '%s': %v", query, err)
	}

	return nil
}
