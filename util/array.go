package util

func InArray[T comparable](value T, array []T) bool {
	for _, i := range array {
		if i == value {
			return true
		}
	}
	return false
}
