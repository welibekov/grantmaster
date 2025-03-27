package utils

import (
	"reflect"
	"sort"

	"github.com/welibekov/grantmaster/internal/policy/types"
)

// Equal compares two slices of Policy and returns true if they are equal.
// It normalizes both slices before comparison to ensure consistent ordering.
func Equal(policyA, policyB []types.Policy) bool {
	// Normalize both policies before comparison to ensure consistency.
	policyA = normalizePolicies(policyA)
	policyB = normalizePolicies(policyB)

	// Use DeepEqual to check if the two normalized policies are identical.
	return reflect.DeepEqual(policyA, policyB)
}

// normalizePolicies takes a slice of Policy, sorts the Roles within each Policy,
// and returns a new slice of Policies with sorted Roles and sorted by Username.
// This ensures consistency in how policies are compared.
func normalizePolicies(policies []types.Policy) []types.Policy {
	sortedPolicies := []types.Policy{}

	// Iterate over each policy to sort its Roles.
	for _, policy := range policies {
		// Sort the Roles slice to ensure a consistent order across policies.
		sort.Strings(policy.Roles)

		// Append the normalized policy to the new slice.
		sortedPolicies = append(sortedPolicies, policy)
	}

	// Sort the policies slice by Username to ensure consistent order.
	sort.Slice(sortedPolicies, func(i, j int) bool {
		return sortedPolicies[i].Username < sortedPolicies[j].Username
	})

	// Return the slice of policies with sorted Roles and order by Username.
	return sortedPolicies
}
