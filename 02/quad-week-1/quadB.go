package piscine

func QuadB(x, y int) {

	cornerA := "/"
	cornerB := "\\"
	chFillerLine := "*"

	cornerC := "\\"
	cornerD := "/"
	chFillerMidle := "*"

	xx := x
	yy := y

	drawer(cornerA, cornerB, chFillerLine, cornerC, cornerD, chFillerMidle, xx, yy)

}
