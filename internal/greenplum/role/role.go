package role

import (
	"context"

	"github.com/welibekov/grantmaster/internal/role/base"
)

// GPRole is a struct that embeds the base.Role type,
// allowing it to inherit its methods and properties.
type GPRole struct {
	*base.Role // Pointer to the base.Role
}

// New creates a new instance of GPRole using the provided context
// and configuration map. It initializes the embedded Role from
// the base package with the provided configuration.
// 
// ctx is the context for the creation process, which may be used
// for cancellation or timeout information.
// cfg is a map of configuration options to initialize the Role.
func New(ctx context.Context, cfg map[string]string) (*GPRole, error) {
	// Create a new GPRole by initializing the embedded Role with the given configuration
	// and return a pointer to the newly created GPRole instance.
	return &GPRole{
		Role: base.NewRole(cfg), // Initialize the embedded Role with the configuration
	}, nil
}
