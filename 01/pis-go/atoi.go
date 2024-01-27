package piscine

func Atoi(str string) int {
	var isNegative bool

	for index, runeValue := range str {
		// check if number has positive symbol
		if runeValue == '+' && index == 0 { // just in first position
			str = str[1:] // remove + from first position
			continue
		}

		// check if negative
		if runeValue == '-' && index == 0 { // just in firist position
			str = str[1:]
			// new string (without first character)
			isNegative = true
			continue
		}

		// Check if the character is a number based on ASCII values
		if !(runeValue >= '0' && runeValue <= '9') {
			return 0 // not valid input
		}
	}

	var result int

	// Convert each digit to int
	for _, runeDigit := range str {
		// Mutiple result by 10 then add the rune decimal rune - 48 (0) to get an absolute value
		result = result*10 + int(runeDigit-'0')
	}

	// Revert number back to negative
	if isNegative {
		result = -result
	}

	return result
}
