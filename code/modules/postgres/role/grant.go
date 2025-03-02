package role

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/welibekov/grantmaster/modules/role/types"
	"github.com/welibekov/grantmaster/modules/template"
)

func (p *PGRole) Grant(ctx context.Context, roles []types.Role) error {
	queryBody, err := template.Generate("postgres/role/grant.tmpl", p.pool, roles)
	if err != nil {
		return fmt.Errorf("couldn't generate grant query template: %v", err)
	}

	query := string(queryBody)

	logrus.Debugln(query) // Log the generated query for debugging purposes

	_, err = p.pool.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("couldn't construct grant query: %v", err)
	}

	return nil
}

// All below functions and methods are not relevant when template generation
// is in place.
//func (p *PGRole) GrantOldVersion(ctx context.Context, roles []types.Role) error {
//	//for _, role := range roles {
//	//	query, err := p.grantQuery(ctx, role)
//	//	if err != nil {
//	//		return fmt.Errorf("couldn't construct grant query: %v", err)
//	//	}
//
//	//	logrus.Debugln(*query) // Log the generated query for debugging purposes
//
//	//	// Execute the grant query on schema in the PostgreSQL database
//	//	_, err = p.pool.Exec(ctx, *query)
//	//	if err != nil {
//	//		// Wrap the error with more context to help identify the issue if it occurs
//	//		return fmt.Errorf("couldn't grants privileges to role '%s': %w", role.Name, err)
//	//	}
//	//}
//
//	// Return nil if all policies were processed without errors
//	return nil
//}
//
//func (p *PGRole) grantQuery(ctx context.Context, role types.Role) (*string, error) {
//	var query string
//
//	exist, err := p.roleExist(ctx, role)
//	if err != nil {
//		return nil, fmt.Errorf("role checking failed: %v", err)
//	}
//
//	if !exist {
//		query += fmt.Sprintf(`CREATE ROLE %s;`, role.Name)
//	}
//
//	for _, schema := range role.Schemas {
//		for _, grant := range schema.Grants {
//			if strings.ToUpper(grant) == "SELECT" {
//				query += fmt.Sprintf(`GRANT %s ON ALL TABLES IN SCHEMA %s TO %s;`, grant, schema.Schema, role.Name)
//			} else {
//				query += fmt.Sprintf(`GRANT %s ON SCHEMA %s TO %s;`, grant, schema.Schema, role.Name)
//			}
//		}
//	}
//
//	return &query, nil
//}
//
//func (p *PGRole) roleExist(ctx context.Context, role types.Role) (bool, error) {
//	query := fmt.Sprintf(`SELECT 1 FROM pg_roles WHERE rolname = '%s';`, role.Name)
//
//	logrus.Debugln(query)
//
//	rows, err := p.pool.Query(ctx, query)
//	if err != nil {
//		return false, fmt.Errorf("couldn't find if role exist: %v", err)
//	}
//	defer rows.Close()
//
//	var exist int
//
//	for rows.Next() {
//		if err := rows.Scan(&exist); err != nil {
//			return false, fmt.Errorf("failed to scan row for checking role: %w", err)
//		}
//
//		break
//	}
//
//	if err := rows.Err(); err != nil {
//		return false, fmt.Errorf("error occurred while iterating over rows: %w", err)
//	}
//
//	return exist == 1, nil
//}
