package gosfunc

func Run[T, R any](it T, block func(it T) R) R {
	return block(it)
}
