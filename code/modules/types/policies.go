package types

// Policy defines a set of permissions for a user,
// specifying which actions they can perform.
type Policy struct {
	Username string   `yaml:"username"` // The username associated with the policy.
	Roles    []string `yaml:"roles"`    // List of roles allowed to perform the action.
}

// Role defines a set of roles that should be created in the system.
type Role struct{}
