package piscine

func UltimateDivMod(a *int, b *int) {
	c := *a
	d := *b
	*a = c / d
	*b = c % d
}
