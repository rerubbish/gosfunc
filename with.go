package gosfunc

func With[T, R any](this T, block func(this T) R) R {
	return block(this)
}
