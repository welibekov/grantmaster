package utils

// InMap checks if a given key exists in a map.
// It takes a map where keys are of type T and values are of type Y, 
// and a key of type T. The function returns true if the key exists in the map, 
// and false otherwise.
//
// T and Y are type parameters, allowing this function to be used with 
// any comparable key type and any value type.
func InMap[T, Y comparable](where map[T]Y, what T) bool {
	_, found := where[what] // Check if the key 'what' exists in the map 'where'
	return found // Return true if found, false otherwise
}
