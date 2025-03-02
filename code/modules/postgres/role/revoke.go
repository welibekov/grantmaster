package role

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/welibekov/grantmaster/modules/role/types"
	"github.com/welibekov/grantmaster/modules/template"
)

func (p *PGRole) Revoke(ctx context.Context, roles []types.Role) error {
	queryBody, err := template.Generate("postgres/role/revoke.tmpl", p.pool, roles)
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

//func (p *PGRole) RevokeOldVersion(ctx context.Context, roles []types.Role) error {
//	for _, role := range roles {
//		query := p.revokeQuery(role)
//
//		logrus.Debugln(query) // Log the generated query for debugging purposes
//
//		// Execute the revoke query on the PostgreSQL database
//		_, err := p.pool.Exec(ctx, query)
//		if err != nil {
//			// Wrap the error with more context to help identify the issue if it occurs
//			return fmt.Errorf("couldn't revoke grants from role'%s': %w", role.Name, err)
//		}
//	}
//
//	// Return nil if all policies were processed without errors
//	return nil
//}
//
//func (p *PGRole) revokeQuery(role types.Role) string {
//	var query string
//	for _, schema := range role.Schemas {
//		query += fmt.Sprintf(`REVOKE ALL PRIVILEGES ON SCHEMA %s FROM %s;`, schema.Schema, role.Name)
//		query += fmt.Sprintf(`REVOKE GRANT OPTION FOR ALL PRIVILEGES ON SCHEMA %s FROM %s;`, schema.Schema, role.Name)
//	}
//
//	//FIXME: Do we need to drop role at the end?
//	//query += fmt.Sprintf(`DROP ROLE %s`, role.Name)
//
//	return query
//}
