package piscine

func Index(s string, toFind string) int {
	if len(toFind) == 0 {
		return 0
	}

	slen := len(s)
	tlen := len(toFind)

	if tlen > slen {
		return -1
	}

	for i := 0; i <= slen-tlen; i++ {
		match := true
		for j := 0; j < tlen; j++ {
			if s[i+j] != toFind[j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}

	return -1
}
