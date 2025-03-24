package role

// TableLevelGrants holds a list of table-level privileges 
// that can be granted to users in a database.
var TableLevelGrants = []string{"SELECT", "INSERT", "UPDATE", "DELETE", "TRUNCATE", "REFERENCES", "TRIGGER", "ALL"}
