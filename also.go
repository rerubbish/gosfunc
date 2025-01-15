package gosfunc

func Also[T any](it T, block func(it T)) (T, error) {
	block(it)
	return it, nil
}
