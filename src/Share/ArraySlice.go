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

// func Any[T any](array []T, predicate func(T) bool) bool {
// 	for _, value := range array {
// 		if predicate(value) {
// 			return true
// 		}
// 	}
// 	return false
// }
