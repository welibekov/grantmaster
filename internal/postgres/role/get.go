package role

import (
	"context"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/welibekov/grantmaster/internal/role/types"
)

// Get retrieves existing roles and their grants from the database.
func (p *PGRole) Get(ctx context.Context) ([]types.Role, error) {
	// rolesMap maps each role to its corresponding schemas and their grants.
	rolesMap := make(map[string]map[string][]string)

	// roles will hold the resulting role objects to be returned.
	roles := make([]types.Role, 0)

	// Construct the SQL query to get roles along with their permissions.
	query := fmt.Sprintf(`
WITH all_roles AS (
  SELECT rolname AS role_or_user
  FROM pg_roles
  WHERE rolname LIKE '%s%%'
),
all_schemas AS (
  SELECT nspname AS schema_name
  FROM pg_namespace
  WHERE nspname NOT LIKE 'pg_%%'
    AND nspname NOT IN ('information_schema', 'public')
),
all_tables AS (
  SELECT c.oid AS table_oid, c.relname, n.nspname AS schema_name
  FROM pg_class c
  JOIN pg_namespace n ON c.relnamespace = n.oid
  WHERE c.relkind IN ('r', 'v', 'm')
    AND n.nspname NOT LIKE 'pg_%%'
    AND n.nspname NOT IN ('information_schema', 'public')
),
schema_permissions AS (
  SELECT 
    r.role_or_user,
    s.schema_name,
    'USAGE' AS permission_name,
    CASE WHEN has_schema_privilege(r.role_or_user, s.schema_name, 'USAGE') THEN 'YES' ELSE 'NO' END AS has_access
  FROM all_roles r
  CROSS JOIN all_schemas s
  UNION ALL
  SELECT 
    r.role_or_user,
    s.schema_name,
    'CREATE' AS permission_name,
    CASE WHEN has_schema_privilege(r.role_or_user, s.schema_name, 'CREATE') THEN 'YES' ELSE 'NO' END AS has_access
  FROM all_roles r
  CROSS JOIN all_schemas s
),
table_permissions AS (
  SELECT 
    r.role_or_user,
    t.schema_name,
    'SELECT' AS permission_name,
    CASE WHEN has_table_privilege(r.role_or_user, t.table_oid, 'SELECT') THEN 'YES' ELSE 'NO' END AS has_access
  FROM all_roles r
  CROSS JOIN all_tables t
  UNION ALL
  SELECT r.role_or_user, t.schema_name, 'INSERT', CASE WHEN has_table_privilege(r.role_or_user, t.table_oid, 'INSERT') THEN 'YES' ELSE 'NO' END
  FROM all_roles r CROSS JOIN all_tables t
  UNION ALL
  SELECT r.role_or_user, t.schema_name, 'UPDATE', CASE WHEN has_table_privilege(r.role_or_user, t.table_oid, 'UPDATE') THEN 'YES' ELSE 'NO' END
  FROM all_roles r CROSS JOIN all_tables t
  UNION ALL
  SELECT r.role_or_user, t.schema_name, 'DELETE', CASE WHEN has_table_privilege(r.role_or_user, t.table_oid, 'DELETE') THEN 'YES' ELSE 'NO' END
  FROM all_roles r CROSS JOIN all_tables t
  UNION ALL
  SELECT r.role_or_user, t.schema_name, 'TRUNCATE', CASE WHEN has_table_privilege(r.role_or_user, t.table_oid, 'TRUNCATE') THEN 'YES' ELSE 'NO' END
  FROM all_roles r CROSS JOIN all_tables t
  UNION ALL
  SELECT r.role_or_user, t.schema_name, 'REFERENCES', CASE WHEN has_table_privilege(r.role_or_user, t.table_oid, 'REFERENCES') THEN 'YES' ELSE 'NO' END
  FROM all_roles r CROSS JOIN all_tables t
  UNION ALL
  SELECT r.role_or_user, t.schema_name, 'TRIGGER', CASE WHEN has_table_privilege(r.role_or_user, t.table_oid, 'TRIGGER') THEN 'YES' ELSE 'NO' END
  FROM all_roles r CROSS JOIN all_tables t
)
SELECT *
FROM (
  SELECT * FROM schema_permissions
  UNION ALL
  SELECT * FROM table_permissions
) all_permissions
ORDER BY role_or_user, schema_name, permission_name
`, p.Prefix)

	logrus.Debugln("role get query:", query)

	// Execute the constructed query against the database.
	rows, err := p.pool.Query(ctx, query)
	if err != nil {
		// Wrap the error to include context about the failure.
		return roles, fmt.Errorf("failed to execute query to get roles: %w", err)
	}
	defer rows.Close() // Ensure rows are closed after processing.

	// Iterate through the result set.
	for rows.Next() {
		var (
			role, schema, permission, hasAccess string
		)

		// Scan the row into the variables for further processing.
		if err := rows.Scan(&role, &schema, &permission, &hasAccess); err != nil {
			// Wrap the error to indicate an issue during scanning.
			return roles, fmt.Errorf("failed to scan row: %w", err)
		}

		// Skip this entry if the role does not have access.
		if hasAccess != "YES" {
			continue
		}

		// Initialize the map entry for the role if it doesn't exist.
		if _, exist := rolesMap[role]; !exist {
			rolesMap[role] = make(map[string][]string)
		}

		// Initialize the map entry for the schema if it doesn't exist.
		if _, found := rolesMap[role][schema]; !found {
			rolesMap[role][schema] = []string{}
		}

		// Append the permission to the list of grants for the schema.
		rolesMap[role][schema] = append(rolesMap[role][schema], strings.ToLower(permission))
	}

	// Check if any errors occurred during the iteration of rows.
	if err := rows.Err(); err != nil {
		// Wrap the error to indicate an issue occurred while processing rows.
		return roles, fmt.Errorf("error occurred while iterating over rows: %w", err)
	}

	// Convert the rolesMap to a list of role objects for output.
	for role, schemasMap := range rolesMap {
		var schemas []types.Schema

		// Iterate over each schema in the map for the role.
		for schema, grants := range schemasMap {
			// Only include schemas that have grants.
			if len(grants) > 0 {
				schemas = append(schemas, types.Schema{
					Schema: schema,
					Grants: grants,
				})
			}
		}

		// Create a Role object and append it to the roles slice.
		roles = append(roles, types.Role{
			Name:    role,
			Schemas: schemas,
		})
	}

	// Return the list of roles with their associated schemas and no error.
	return roles, nil
}
