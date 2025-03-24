package role

import (
	"context"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/welibekov/grantmaster/modules/role/types"
	"github.com/welibekov/grantmaster/modules/utils"
)

// IsExist checks if a given query returns any results.
// It returns true if at least one row is found, otherwise returns false.
func (p *PGRole) IsExist(ctx context.Context, query string) (bool, error) {
	logrus.Debugln(query)

	// Execute the query using the connection pool.
	rows, err := p.pool.Query(ctx, query)
	if err != nil {
		return false, fmt.Errorf("couldn't run query: %v", err)
	}
	defer rows.Close() // Ensure that rows are closed after we're done with them.

	var exist int

	// Iterate through the result set.
	for rows.Next() {
		// Scan the result into the exist variable.
		if err := rows.Scan(&exist); err != nil {
			return false, fmt.Errorf("failed to scan row for 'exist' value: %v", err)
		}

		// Since we're only checking for existence, we can break after the first row.
		break
	}

	// Check if there was an error during iteration.
	if err := rows.Err(); err != nil {
		return false, fmt.Errorf("error occurred while iterating over rows: %v", err)
	}

	// Return true if `exist` equals 1, indicating the row was found.
	return exist == 1, nil
}

// IsRoleExist checks if a role with the given name exists in the database.
// It constructs a query to check for the role and calls IsExist method.
func (p *PGRole) IsRoleExist(ctx context.Context, role types.Role) (bool, error) {
	query := fmt.Sprintf(`SELECT 1 FROM pg_roles WHERE rolname = '%s';`, role.Name)

	return p.IsExist(ctx, query)
}

// IsItTableLevelGrant determines if the provided grant is a table-level grant.
// It checks against a predefined list of table-level grants.
func (p *PGRole) IsItTableLevelGrant(grant string) bool {
	return utils.In(strings.ToUpper(grant), TableLevelGrants)
}

// IsTablesExistInSchema checks if there are tables in the given schema.
// It returns true if tables exist, or false if there are none or if an error occurs.
func (p *PGRole) IsTablesExistInSchema(ctx context.Context, schema types.Schema) bool {
    query := fmt.Sprintf(`
SELECT COUNT(*) 
FROM information_schema.tables
WHERE table_schema = '%s' AND table_type = 'BASE TABLE';
`, schema.Schema)

	// Call IsExist to determine if tables exist and handle any errors reported.
	exist, err := p.IsExist(ctx, query)
	if err != nil {
		logrus.Warnf("checking tables existence in schema %s failed: %v", schema.Schema, err)
		return false
	}

	return exist // Return the result of existence check.
}
