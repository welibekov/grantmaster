package policy

import (
	"context"

	"github.com/welibekov/grantmaster/modules/policy/base"
)

// GPPolicy is a struct that embeds the base.Policy struct,
// extending its functionality for specific policy needs.
type GPPolicy struct {
	*base.Policy // Embedding base.Policy to inherit its methods and fields
}

// New creates a new instance of GPPolicy. It takes a context and a configuration map
// as arguments and initializes a new base.Policy using the provided configuration.
func New(ctx context.Context, cfg map[string]string) (*GPPolicy, error) {
	// Initialize the GPPolicy with a base.Policy created from the provided configuration
	return &GPPolicy{
		Policy: base.NewPolicy(cfg),
	}, nil // Return the new GPPolicy instance and a nil error
}
