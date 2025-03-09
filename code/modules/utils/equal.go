package utils

import (
	"bytes"
	"encoding/json"
)

// Equal marshals two interfaces to []byte and compares them.
// It returns true if they are equal, otherwise false.
func Equal(slice1, slice2 interface{}) (bool, error) {
	// Marshal the first slice
	data1, err1 := json.Marshal(slice1)
	if err1 != nil {
		return false, err1
	}

	// Marshal the second slice
	data2, err2 := json.Marshal(slice2)
	if err2 != nil {
		return false, err2
	}

	// Compare the byte slices
	return bytes.Equal(data1, data2), nil
}
