package array

func Map[T interface{}, R interface{}](array []T, callback func(T) R) []R {
	var result []R
	for _, v := range array {
		result = append(result, callback(v))
	}
	return result
}

func Find[T interface{}](array []T, callback func(T) bool) *T {
	for _, v := range array {
		if callback(v) {
			return &v
		}
	}
	return nil
}

func Filter[T interface{}](array []T, callback func(T) bool) []T {
	var result []T
	for _, v := range array {
		if callback(v) {
			result = append(result, v)
		}
	}
	return result
}
