package piscine

func UltimateDivMod(a *int, b *int) {
	acopy := *a
	bcopy := *b
	*a = acopy / bcopy
	*b = acopy % bcopy
}
