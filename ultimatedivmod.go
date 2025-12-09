package student

func UltimateDivMod(a *int, b *int) {
	if *b == 0 {
		return
	}
	div := *a / *b
	mod := *a % *b
	*a = div
	*b = mod
}
