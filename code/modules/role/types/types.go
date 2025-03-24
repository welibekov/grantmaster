package types

// Role defines a set of roles that should be created in the system.
type Role struct {
	Name    string   `yaml:"name"`    // The name associated with the role (e.g., "Admin", "User")
	Schemas []Schema `yaml:"schemas"` // A slice of Schema objects associated with the role
}

// Schema represents a database schema and the associated grants for that schema.
type Schema struct {
	Schema string   `yaml:"schema"` // The name of the schema (e.g., "public", "sales_data")
	Grants []string `yaml:"grants"` // A list of grants (e.g., "SELECT", "INSERT") that this role has on the schema
}
