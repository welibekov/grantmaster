package role

import "fmt"

// getQuery constructs and returns a SQL query that retrieves the permissions 
// of roles or users in the PostgreSQL database. The query checks for schema 
// usage and various table permissions (such as SELECT, INSERT, UPDATE, 
// DELETE, TRUNCATE, REFERENCES, and TRIGGER) for roles that match a given 
// prefix. It utilizes Common Table Expressions (CTEs) to gather roles, 
// schemas, and tables, and combines the results into a single list of 
// permissions, ordered by role or user, schema name, and permission name.
func (p *PGRole) getQuery() string {
	return fmt.Sprintf(`
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
}
