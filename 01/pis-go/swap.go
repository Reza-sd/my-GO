package piscine

func Swap(a *int, b *int) {
	aa := *a
	bb := *b

	*a = bb
	*b = aa
}
