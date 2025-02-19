package types

// Action represents a specific action that can be performed,
// along with the roles that are permitted to perform it.
type Action struct {
	Action string   `yaml:"action"` // The name of the action.
	Roles  []string `yaml:"roles"`  // List of roles allowed to perform the action.
}

// Policy defines a set of permissions for a user,
// specifying which actions they can perform.
type Policy struct {
	Username string   `yaml:"username"` // The username associated with the policy.
	Actions  []Action `yaml:"actions"`  // List of actions permitted for the user.
}
