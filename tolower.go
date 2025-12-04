package piscine

func ToLower(s string) string {
	text := []rune(s)
	for i, r := range text {
		if r >= 'A' && r <= 'Z' {
			text[i] = r + ('a' - 'A')
		}
	}
	return string(text)
}
