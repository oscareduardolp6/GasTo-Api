package share

func RemoveFromArrayUnsafe[T any](array []T, index int) []T {
	return append(array[:index], array[index+1:]...)
}
