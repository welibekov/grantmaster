package utils

import (
	"bytes"
	"encoding/json"
)

// Equal marshals two interfaces to []byte and compares them.
// It returns true if they are equal, otherwise false and any error encountered during marshaling.
func Equal(slice1, slice2 interface{}) (bool, error) {
	// Marshal the first slice into JSON format
	data1, err1 := json.Marshal(slice1)
	if err1 != nil {
		// Return false and the error if marshaling fails
		return false, err1
	}

	// Marshal the second slice into JSON format
	data2, err2 := json.Marshal(slice2)
	if err2 != nil {
		// Return false and the error if marshaling fails
		return false, err2
	}

	// Compare the byte slices representing the JSON-encoded data
	return bytes.Equal(data1, data2), nil
}
