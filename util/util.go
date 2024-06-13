package util

func AsPtr[T any](value T) *T {
	return &value
}

func Ternary[T any](condition bool, trueValue T, falseValue T) T {
	if condition {
		return trueValue
	}
	return falseValue
}
