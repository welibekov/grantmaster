package role

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/welibekov/grantmaster/modules/role/types"
	"github.com/welibekov/grantmaster/modules/template"
)

func (p *PGRole) Revoke(ctx context.Context, roles []types.Role) error {
	tmpl, err := template.New("postgres/role/revoke.tmpl", p.pool)
	if err != nil {
		return fmt.Errorf("couldn't create template: %v", err)
	}

	queryBody, err := tmpl.Generate(ctx, roles)
	if err != nil {
		return fmt.Errorf("couldn't generate revoke query template: %v", err)
	}

	query := string(queryBody)

	logrus.Debugln(query) // Log the generated query for debugging purposes

	_, err = p.pool.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("couldn't construct grant query: %v", err)
	}

	return nil
}
