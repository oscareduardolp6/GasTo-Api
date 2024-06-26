package share

type Criteria[T any] interface {
	Filter(val T) bool
	SortingLess(val1, val2 T) bool
}
