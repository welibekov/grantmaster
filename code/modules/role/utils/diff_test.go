package utils

import (
	"log"
	"testing"

	"github.com/welibekov/grantmaster/modules/role/types"
	"github.com/welibekov/grantmaster/modules/utils"
)

// TestDiff tests the Diff function which is the public interface.
func TestDiff(t *testing.T) {
	// Define test cases for diffing roles.
	tests := []struct {
		name     string      // test case name
		roles    []types.Role // original roles
		input    []types.Role // roles to compare against
		expected []types.Role // expected result of diff operation
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

	// Run each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Diff(tt.roles, tt.input)

			// If both result and expected are nil or empty, they are considered equal.
			if result == nil && len(tt.expected) == 0 {
				return
			}

			// Compare the result with the expected output.
			equal, err := utils.Equal(result, tt.expected)
			if err != nil {
				log.Fatal("couldn't compare to slices")
			}

			// If not equal, report the difference.
			if !equal {
				t.Errorf("Diff() = %v; want %v", result, tt.expected)
			}
		})
	}
}

// TestDiffRoles tests the diffRoles function directly.
func TestDiffRoles(t *testing.T) {
	// Simulating test data for role comparisons
	tests := []struct {
		name     string      // test case name
		first    []types.Role // first slice of roles
		second   []types.Role // second slice of roles for comparison
		expected []types.Role // expected result of the diff operation
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

	// Run each test case.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := diffRoles(tt.first, tt.second)

			// If both result and expected are nil or empty, they are considered equal.
			if result == nil && len(tt.expected) == 0 {
				return
			}

			// Compare the result with the expected output.
			equal, err := utils.Equal(result, tt.expected)
			if err != nil {
				log.Fatal("couldn't compare to slices")
			}

			// If not equal, report the difference.
			if !equal {
				t.Errorf("diffRoles() = %v; want %v", result, tt.expected)
			}
		})
	}
}

// TestDiffSchemas tests the diffSchemas function directly.
func TestDiffSchemas(t *testing.T) {
	// Define test cases for diffing schemas.
	tests := []struct {
		name     string       // test case name
		first    []types.Schema // first slice of schemas
		second   []types.Schema // second slice of schemas for comparison
		expected []types.Schema // expected result of the diff operation
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

	// Run each test case.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := diffSchemas(tt.first, tt.second)

			// If both result and expected are nil or empty, they are considered equal.
			if result == nil && len(tt.expected) == 0 {
				return
			}

			// Compare the result with the expected output.
			equal, err := utils.Equal(result, tt.expected)
			if err != nil {
				log.Fatal("couldn't compare to slices")
			}

			// If not equal, report the difference.
			if !equal {
				t.Errorf("diffSchemas() = %v; want %v", result, tt.expected)
			}
		})
	}
}

// TestDiffGrants tests the diffGrants function directly.
func TestDiffGrants(t *testing.T) {
	// Define test cases for diffing grants.
	tests := []struct {
		name     string   // test case name
		first    []string // first slice of grants
		second   []string // second slice of grants for comparison
		expected []string // expected result of the diff operation
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

	// Run each test case.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := diffGrants(tt.first, tt.second)

			// If both result and expected are nil or empty, they are considered equal.
			if result == nil && len(tt.expected) == 0 {
				return
			}

			// Compare the result with the expected output.
			equal, err := utils.Equal(result, tt.expected)
			if err != nil {
				log.Fatal("couldn't compare to slices")
			}

			// If not equal, report the difference.
			if !equal {
				t.Errorf("diffGrants() = %v; want %v", result, tt.expected)
			}
		})
	}
}
