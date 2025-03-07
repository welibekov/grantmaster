package role

import (
	"context"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/welibekov/grantmaster/modules/role/types"
	"github.com/welibekov/grantmaster/modules/utils"
)

func (p *PGRole) Revoke(ctx context.Context, roles []types.Role) error {
	var query string

	for _, role := range roles {
		exist, err := p.IsRoleExist(ctx, role)
		if err != nil {
			return fmt.Errorf("couldn't check if role %s exist: %v", role.Name, err)
		}

		if exist {
			for _, schema := range role.Schemas {
				for _, grant := range schema.Grants {
					if utils.In(strings.ToLower(grant), []string{"select"}) {
						query += fmt.Sprintf("REVOKE %s ON ALL TABLES IN SCHEMA %s FROM %s;", grant, schema.Schema, role.Name)
					} else {
						query += fmt.Sprintf("REVOKE %s ON SCHEMA %s FROM %s;", grant, schema.Schema, role.Name)
					}
				}
			}
		}
	}

	logrus.Debugln(query) // Log the generated query for debugging purposes

	_, err := p.pool.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("couldn't execute drop query '%s': %v", query, err)
	}

	return nil
}
