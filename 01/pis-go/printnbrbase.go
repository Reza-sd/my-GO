package piscine

func PrintNbrBase(nbr int, base string) {
	nbase := len(base)
	vbase := []int{}
	for i := 1; i <= nbase; i++ {
		vbase = append(vbase, i)
	}
	reminder := 0
	remindersArr := []int{}
	num := nbr
	for {
		reminder = num % nbase
		remindersArr = append(remindersArr, reminder)
		num = num / nbase
	}
}
