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
    -- Explicitly checking USAGE privilege for schemas.
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

    -- Explicitly checking CREATE privilege for schemas.
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

	// Execute the constructed query against the database.
	rows, err := p.pool.Query(ctx, query)
	if err != nil {
		// Wrap the error to include context about the failure.
		return roles, fmt.Errorf("failed to execute query to get policies: %w", err)
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
