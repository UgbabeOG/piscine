package piscine

func ConcatParams(args []string) string {
	var result string

	for i, str := range args {
		if i > 0 {
			result += "\n"
		}
		result += str
	}

	return result
}
