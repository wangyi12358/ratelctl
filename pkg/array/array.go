package array

func Map[T interface{}, V interface{}](array []T, callback func(T) V) []V {
	var result []V
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
