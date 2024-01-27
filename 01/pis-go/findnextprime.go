package piscine

func FindNextPrime(nb int) int {
	if IsPrime(nb) == true {
		return nb
	} else {
		for {

			nb++
			if IsPrime(nb) == true {
				return nb
			}
		}
	}
}
