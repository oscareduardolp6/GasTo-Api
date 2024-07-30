package share

func RemoveFromArrayUnsafe[T any](array []T, index int) []T {
	return append(array[:index], array[index+1:]...)
}

func Any[T any](predicate func(T) bool, values ...T) bool {
	for _, value := range values {
		if predicate(value) {
			return true
		}
	}
	return false
}

func Filter[T any](predicate func(T) bool, values ...T) []T {
	filteredValues := make([]T, 0)
	for _, value := range values {
		if predicate(value) {
			filteredValues = append(filteredValues, value)
		}
	}
	return filteredValues
}

func Find[T any](predicate func(T) bool, values []T) *T {
	for _, value := range values {
		if predicate(value) {
			result := value
			return &result
		}
	}
	return nil
}
