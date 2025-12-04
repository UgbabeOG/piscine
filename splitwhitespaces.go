package piscine

func SplitWhiteSpaces(s string) []string {
	var words []string
	current := ""

	for _, r := range s {
		if r == ' ' || r == '\t' || r == '\n' {
			if current != "" {
				words = append(words, current)
				current = ""
			}
		} else {
			current += string(r)
		}
	}

	if current != "" {
		words = append(words, current)
	}

	return words
}
