package types

// Role defines a set of roles that should be created in the system.
type Role struct {
	Name   string `yaml:"name"`   // The name associated with the role
	Type   string `yaml:"type"`   // Type of the role.
	Schema string `yaml:"schema"` // Schema name associated with the role
}
