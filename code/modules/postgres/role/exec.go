package role

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/welibekov/grantmaster/modules/role/types"
	"github.com/welibekov/grantmaster/modules/template"
)

func (p *PGRole) exec(ctx context.Context, roles []types.Role, path string) error {
	tmpl, err := template.New(path, p.pool)
	if err != nil {
		return fmt.Errorf("couldn't create new template: %v", err)
	}

	queryBody, err := tmpl.Generate(ctx, roles)
	if err != nil {
		return fmt.Errorf("couldn't generate query template '%s': %v", path, err)
	}

	query := string(queryBody)

	logrus.Debugln(query) // Log the generated query for debugging purposes

	_, err = p.pool.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("couldn't construct grant query '%s': %v", path, err)
	}

	return nil
}
