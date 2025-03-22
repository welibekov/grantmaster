package utils

import (
	"reflect"
	"sort"

	"github.com/welibekov/grantmaster/modules/policy/types"
)

// Equal compares two slices of Policy and returns true if they are equal.
func Equal(policyA, policyB []types.Policy) bool {
	// Normalize both policies before comparison to ensure consistency.
	policyA = normalizePolicies(policyA)
	policyB = normalizePolicies(policyB)

	// Use DeepEqual to check if the two normalized policies are the same.
	return reflect.DeepEqual(policyA, policyB)
}

// normalizePolicies takes a slice of Policy, sorts the Roles within each Policy,
// and returns a new slice of Policies with sorted Roles.
func normalizePolicies(policies []types.Policy) []types.Policy {
	sortedPolicies := []types.Policy{}

	// Iterate over each policy to sort its Roles.
	for _, policy := range policies {
		// Sort the Roles slice to ensure a consistent order.
		sort.Strings(policy.Roles)

		// Append the normalized policy to the new slice.
		sortedPolicies = append(sortedPolicies, policy)
	}

	// Return the slice of policies with sorted Roles.
	return sortedPolicies
}
