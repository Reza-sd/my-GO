package piscine

func ConcatParams(args []string) string {
	// make
	str := ""
	la := len(args)
	for index, item := range args {
		if index < la-1 {
			str += item + "\n"
		} else if index == la-1 {
			str += item
		}
	}
	return str
}
