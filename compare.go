package piscine

func Compare(a, b string) int {
	alen := len(a)
	blen := len(b)

	for i := 0; i < alen && i < blen; i++ {
		if a[i] != b[i] {
			if a[i] < b[i] {
				return -1
			}
			return 1
		}
	}

	if alen < blen {
		return -1
	} else if alen > blen {
		return 1
	}

	return 0
}
