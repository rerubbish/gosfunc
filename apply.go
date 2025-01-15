package gosfunc

func Apply[T any](this T, block func(this T)) T {
	block(this)
	return this
}
