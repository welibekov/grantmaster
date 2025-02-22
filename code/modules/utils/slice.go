package utils

func In[T comparable](what T, where []T) bool {
	for _, item := range where {
		if item == what {
			return true
		}
	}

	return false
}
