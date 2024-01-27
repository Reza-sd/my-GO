package piscine

// -----------math power function--------------------------
func power(a, b int) int {
	if b == 0 {
		return 1
	}
	r := 1
	for i := 1; i <= b; i++ {
		r = r * a
	}
	return r
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	// ==========================baseForm to Decimal=totalValueNbr-====================
	//---------setup------------

	runeNbr := []rune(nbr)
	lrn := len(runeNbr)
	runeBf := []rune(baseFrom)
	lbf := len(runeBf) //=base from number
	powerTo := 0       // value * (base^powerTo)
	totalValueNbr := 0 // totalValueNbr = totalValueNbr + (value of each letter)

	//----------calculate totalValueNbr=>Decimal-----------

	for iRnbr := lrn - 1; iRnbr >= 0; iRnbr-- {
		for index, v := range baseFrom {
			if v == runeNbr[iRnbr] {
				totalValueNbr += power(lbf, powerTo) * index
			}
		}
		powerTo++
	}

	//=======================Decimal=totalValueNbr to baseTo-==========================
	//---------setup---------------

	runeBT := []rune(baseTo)
	var valueArr []int
	var reminder int
	lrBT := len(runeBT)

	//-------calculating values and save it in a slice(array)---------

	for {
		reminder = totalValueNbr % lrBT
		valueArr = append(valueArr, reminder)
		totalValueNbr = totalValueNbr / lrBT
		if totalValueNbr == 0 {
			break
		}
	}

	//--------reverse the slice and find the character incharge with its value---------------
	var resultRune []rune
	for i := len(valueArr) - 1; i >= 0; i-- {
		resultRune = append(resultRune, runeBT[valueArr[i]])
	}
	//-----------------return--------------------------
	return string(resultRune)
}

/*
//=====================test-case================================
	fmt.Println(piscine.ConvertBase("101011", "01", "0123456789")) // 43
	fmt.Println(piscine.ConvertBase("125", "0123456789", "01")) // 1111101
	fmt.Println(piscine.ConvertBase("125", "0123456789", "0123456789ABCDEF")) // 7D
	fmt.Println(piscine.ConvertBase("101011", "01", "0123456789ABCDEF")) // 2B
	fmt.Println(piscine.ConvertBase("E2", "0123456789ABCDEF", "01")) //11100010
//========================================================
*/
