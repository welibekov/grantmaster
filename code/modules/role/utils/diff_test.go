package utils

import (
	"log"
	"testing"

	"github.com/welibekov/grantmaster/modules/role/types"
	"github.com/welibekov/grantmaster/modules/utils"
)

// TestDiff tests the Diff function which is the public interface.
func TestDiff(t *testing.T) {
	tests := []struct {
		name     string
		roles    []types.Role
		input    []types.Role
		expected []types.Role
	}{
		{
			name: "No differences",
			roles: []types.Role{
				{Name: "Admin", Schemas: []types.Schema{{Schema: "users", Grants: []string{"create", "read"}}}},
			},
			input: []types.Role{
				{Name: "Admin", Schemas: []types.Schema{{Schema: "users", Grants: []string{"create", "read"}}}},
			},
			expected: []types.Role{},
		},
		{
			name: "Role missing in second slice",
			roles: []types.Role{
				{Name: "Admin", Schemas: []types.Schema{{Schema: "users", Grants: []string{"create", "read"}}}},
			},
			input: []types.Role{
				{Name: "User", Schemas: []types.Schema{{Schema: "users", Grants: []string{"read"}}}},
			},
			expected: []types.Role{
				{Name: "Admin", Schemas: []types.Schema{{Schema: "users", Grants: []string{"create", "read"}}}},
			},
		},
		{
			name: "Differences in schemas",
			roles: []types.Role{
				{Name: "Admin", Schemas: []types.Schema{{Schema: "users", Grants: []string{"create"}}}},
			},
			input: []types.Role{
				{Name: "Admin", Schemas: []types.Schema{{Schema: "users", Grants: []string{"read"}}}},
			},
			expected: []types.Role{
				{Name: "Admin", Schemas: []types.Schema{{Schema: "users", Grants: []string{"create"}}}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Diff(tt.roles, tt.input)

			if result == nil && len(tt.expected) == 0 {
				return
			}

			equal, err := utils.Equal(result, tt.expected)
			if err != nil {
				log.Fatal("couldn't compare to slices")
			}

			if !equal {
				t.Errorf("Diff() = %v; want %v", result, tt.expected)
			}
		})
	}
}

// TestDiffRoles tests the diffRoles function directly.
func TestDiffRoles(t *testing.T) {
	// Simulating test data
	tests := []struct {
		name     string
		first    []types.Role
		second   []types.Role
		expected []types.Role
	}{
		{
			name: "Simply missing roles",
			first: []types.Role{
				{Name: "Admin", Schemas: []types.Schema{{Schema: "users", Grants: []string{"create"}}}},
				{Name: "Editor", Schemas: []types.Schema{{Schema: "posts", Grants: []string{"update"}}}},
			},
			second: []types.Role{
				{Name: "User", Schemas: []types.Schema{{Schema: "users", Grants: []string{"read"}}}},
			},
			expected: []types.Role{
				{Name: "Admin", Schemas: []types.Schema{{Schema: "users", Grants: []string{"create"}}}},
				{Name: "Editor", Schemas: []types.Schema{{Schema: "posts", Grants: []string{"update"}}}},
			},
		},
		{
			name: "Schema differences",
			first: []types.Role{
				{Name: "Admin", Schemas: []types.Schema{{Schema: "users", Grants: []string{"read", "write"}}}},
			},
			second: []types.Role{
				{Name: "Admin", Schemas: []types.Schema{{Schema: "users", Grants: []string{"read"}}}},
			},
			expected: []types.Role{
				{Name: "Admin", Schemas: []types.Schema{{Schema: "users", Grants: []string{"write"}}}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := diffRoles(tt.first, tt.second)

			if result == nil && len(tt.expected) == 0 {
				return
			}

			equal, err := utils.Equal(result, tt.expected)
			if err != nil {
				log.Fatal("couldn't compare to slices")
			}

			if !equal {
				t.Errorf("diffRoles() = %v; want %v", result, tt.expected)
			}
		})
	}
}

// TestDiffSchemas tests the diffSchemas function directly.
func TestDiffSchemas(t *testing.T) {
	tests := []struct {
		name     string
		first    []types.Schema
		second   []types.Schema
		expected []types.Schema
	}{
		{
			name: "Simply missing schemas",
			first: []types.Schema{
				{Schema: "users", Grants: []string{"create"}},
				{Schema: "posts", Grants: []string{"read", "update"}},
			},
			second: []types.Schema{
				{Schema: "comments", Grants: []string{"read"}},
			},
			expected: []types.Schema{
				{Schema: "users", Grants: []string{"create"}},
				{Schema: "posts", Grants: []string{"read", "update"}},
			},
		},
		{
			name: "Grants differ in schemas",
			first: []types.Schema{
				{Schema: "users", Grants: []string{"create", "read"}},
			},
			second: []types.Schema{
				{Schema: "users", Grants: []string{"read"}},
			},
			expected: []types.Schema{
				{Schema: "users", Grants: []string{"create"}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := diffSchemas(tt.first, tt.second)

			if result == nil && len(tt.expected) == 0 {
				return
			}

			equal, err := utils.Equal(result, tt.expected)
			if err != nil {
				log.Fatal("couldn't compare to slices")
			}

			if !equal {
				t.Errorf("diffSchemas() = %v; want %v", result, tt.expected)
			}
		})
	}
}

// TestDiffGrants tests the diffGrants function directly.
func TestDiffGrants(t *testing.T) {
	tests := []struct {
		name     string
		first    []string
		second   []string
		expected []string
	}{
		{
			name:     "Missing grants",
			first:    []string{"create", "read"},
			second:   []string{"read"},
			expected: []string{"create"},
		},
		{
			name:     "All grants present",
			first:    []string{"create", "read"},
			second:   []string{"create", "read"},
			expected: []string{},
		},
		{
			name:     "Extra grants in first",
			first:    []string{"create", "delete"},
			second:   []string{"create"},
			expected: []string{"delete"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := diffGrants(tt.first, tt.second)

			if result == nil && len(tt.expected) == 0 {
				return
			}

			equal, err := utils.Equal(result, tt.expected)
			if err != nil {
				log.Fatal("couldn't compare to slices")
			}

			if !equal {
				t.Errorf("diffGrants() = %v; want %v", result, tt.expected)
			}
		})
	}
}
