package role

import (
	"context"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/welibekov/grantmaster/modules/role/types"
)

// Get retrieves existing roles and their grants from the database.
func (p *PGRole) Get(ctx context.Context) ([]types.Role, error) {
	// Map to hold schemas and their grants associated with each role.
	rolesMap := make(map[string]map[string][]string)

	// Slice to hold the resulting roles.
	roles := make([]types.Role, 0)

	// Get roles with specific role prefix.
	query := fmt.Sprintf(`
WITH granted_table_permissions AS (
    SELECT 
        grantee AS role_or_user,
        table_schema AS schema_name,
        privilege_type AS permission_name,
        'YES' AS has_access
    FROM information_schema.role_table_grants
    WHERE table_schema NOT LIKE 'pg_%%'
    AND table_schema NOT IN ('information_schema', 'public')
    AND grantee LIKE '%s%%'
),
all_table_permissions AS (
    SELECT 
        r.rolname AS role_or_user,
        n.nspname AS schema_name,
        unnest(ARRAY['SELECT', 'INSERT', 'UPDATE', 'DELETE', 'TRUNCATE', 'REFERENCES', 'TRIGGER']) AS permission_name,
        'NO' AS has_access
    FROM pg_class c
    JOIN pg_namespace n ON c.relnamespace = n.oid
    CROSS JOIN pg_roles r
    WHERE c.relkind IN ('r', 'v', 'm')  
    AND n.nspname NOT LIKE 'pg_%%'
    AND n.nspname NOT IN ('information_schema', 'public')
    AND r.rolname LIKE '%s%%'
),
granted_schema_permissions AS (
    -- Explicitly checking USAGE privilege
    SELECT 
        r.rolname AS role_or_user,
        n.nspname AS schema_name,
        'USAGE' AS permission_name,
        CASE 
            WHEN has_schema_privilege(r.rolname, n.oid, 'USAGE') THEN 'YES' 
            ELSE 'NO' 
        END AS has_access
    FROM pg_namespace n
    CROSS JOIN pg_roles r
    WHERE n.nspname NOT LIKE 'pg_%%'
    AND n.nspname NOT IN ('information_schema', 'public')
    AND r.rolname LIKE '%s%%'

    UNION ALL

    -- Explicitly checking CREATE privilege
    SELECT 
        r.rolname AS role_or_user,
        n.nspname AS schema_name,
        'CREATE' AS permission_name,
        CASE 
            WHEN has_schema_privilege(r.rolname, n.oid, 'CREATE') THEN 'YES' 
            ELSE 'NO' 
        END AS has_access
    FROM pg_namespace n
    CROSS JOIN pg_roles r
    WHERE n.nspname NOT LIKE 'pg_%%'
    AND n.nspname NOT IN ('information_schema', 'public')
    AND r.rolname LIKE '%s%%'
),
all_schema_permissions AS (
    SELECT 
        r.rolname AS role_or_user,
        n.nspname AS schema_name,
        unnest(ARRAY['USAGE', 'CREATE']) AS permission_name,
        'NO' AS has_access
    FROM pg_namespace n
    CROSS JOIN pg_roles r
    WHERE n.nspname NOT LIKE 'pg_%%'
    AND n.nspname NOT IN ('information_schema', 'public')
    AND r.rolname LIKE '%s%%'
)
SELECT 
    ap.role_or_user,
    ap.schema_name,
    ap.permission_name,
    COALESCE(gp.has_access, 'NO') AS has_access
FROM (
    SELECT role_or_user, schema_name, permission_name, has_access FROM all_table_permissions
    UNION ALL
    SELECT role_or_user, schema_name, permission_name, has_access FROM all_schema_permissions
) ap
LEFT JOIN (
    SELECT role_or_user, schema_name, permission_name, has_access FROM granted_table_permissions
    UNION ALL
    SELECT role_or_user, schema_name, permission_name, has_access FROM granted_schema_permissions
) gp
    ON ap.role_or_user = gp.role_or_user
    AND ap.schema_name = gp.schema_name
    AND ap.permission_name = gp.permission_name
ORDER BY ap.role_or_user, ap.schema_name, ap.permission_name;
`, p.Prefix, p.Prefix, p.Prefix, p.Prefix, p.Prefix)

	logrus.Debugln("role get query:", query)

	rows, err := p.pool.Query(ctx, query)
	if err != nil {
		// Wrap the error to include context about where it occurred.
		return roles, fmt.Errorf("failed to execute query to get policies: %w", err)
	}
	defer rows.Close() // Ensure rows are closed after processing.

	// Process each row returned by the query.
	for rows.Next() {
		var (
			role, schema, permission, hasAccess string
		)

		// Scan the row into the username and role variables.
		if err := rows.Scan(&role, &schema, &permission, &hasAccess); err != nil {
			// Wrap the error to include context about scanning the row.
			return roles, fmt.Errorf("failed to scan row: %w", err)
		}

		if hasAccess != "YES" { // if no grant provided just continue.
			continue
		}

		if _, exist := rolesMap[role]; !exist { // if no role exist in map just create it.
			rolesMap[role] = make(map[string][]string)
		}

		if _, found := rolesMap[role][schema]; !found { // if no schema found in map create it.
			rolesMap[role][schema] = []string{}
		}

		// add grants to the map.
		rolesMap[role][schema] = append(rolesMap[role][schema], strings.ToLower(permission))
	}

	// Check for errors that may have occurred during the row iteration.
	if err := rows.Err(); err != nil {
		// Wrap the error to indicate an issue occurred while processing rows.
		return roles, fmt.Errorf("error occurred while iterating over rows: %w", err)
	}

	// Convert the rolesMap to a list of policies.
	for role, schemasMap := range rolesMap {
		var schemas []types.Schema

		for schema, grants := range schemasMap {
			if len(grants) > 0 {
				schemas = append(schemas, types.Schema{
					Schema: schema,
					Grants: grants,
				})
			}
		}

		roles = append(roles, types.Role{
			Name:    role,
			Schemas: schemas,
		})
	}

	// Return the list of policies and no error.
	return roles, nil
}
