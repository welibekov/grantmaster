package utils

// In checks if the specified element (what) is present in the provided slice (where).
// It uses generics to allow for any type that supports comparison (comparable).
//
// Parameters:
// - what: the element to search for in the slice
// - where: the slice of elements to search within
//
// Returns:
// - true if the element is found in the slice, false otherwise.
func In[T comparable](what T, where []T) bool {
	// Iterate over each item in the slice
	for _, item := range where {
		// Check if the current item is equal to the element we are searching for
		if item == what {
			return true // Return true if we found a match
		}
	}

	// If we finish the loop without finding a match, return false
	return false // Element not found in the slice
}
