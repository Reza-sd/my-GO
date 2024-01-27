package piscine

import "fmt"

func drawer(cornerA, cornerB, chFillerLine, cornerC, cornerD, chFillerMidle string, x, y int) {

	//----------------first-line------------------------------

	dash := ""

	line1 := ""

	for i := 0; i < x-2; i++ {

		dash += chFillerLine

	}

	if x == 1 {

		line1 = cornerA

	} else if x > 1 {

		line1 = cornerA + dash + cornerB

	}

	//--------------last-line-------------------
	lastline := ""
	dash = ""

	for i := 0; i < x-2; i++ {

		dash += chFillerLine

	}

	if x == 1 {

		lastline = cornerC

	} else if x > 1 {

		lastline = cornerC + dash + cornerD

	}

	//--------------midle-lines-------------------

	pipelines := ""

	spaceNum := x - 2

	spaceline := ""

	for i := 0; i < spaceNum; i++ {

		spaceline += " "

	}

	for i := 0; i < y-2; i++ {

		if x == 1 {
			pipelines += chFillerMidle

		} else {
			pipelines += chFillerMidle + spaceline + chFillerMidle
		}

		if !(i == y-3) {
			pipelines += "\n"
		}

	}
	//==============Prints==============================
	//----------------------exeption x=1--------
	//--*-------------exception--1,3-------
	if x == 1 && y == 3 {
		pipelines = chFillerMidle //"|"
	}

	//----------------print fist line-----------------
	if y > 0 && x > 0 {

		fmt.Println(line1) // (4,0)

	}

	//--------------print-midle-lines--------------------
	if y > 2 {

		fmt.Println(pipelines)

	}
	//---------------------print-lastline-------
	if y > 1 {

		fmt.Println(lastline)

	}

	//-------------------------------------

}
