package piscine

func Split(s, sep string) []string {
	var res []string
	start := 0

	for i := 0; i <= len(s); i++ {
		if i == len(s) || (i+len(sep) <= len(s) && s[i:i+len(sep)] == sep) {
			res = append(res, s[start:i])
			if i == len(s) {
				break
			}
			start = i + len(sep)
			i = start - 1
		}
	}
	return res
}
