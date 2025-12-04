package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	// Convert from baseFrom to decimal
	n := 0
	for _, r := range nbr {
		val := 0
		for i, b := range baseFrom {
			if b == r {
				val = i
				break
			}
		}
		n = n*len(baseFrom) + val
	}

	// Handle zero
	if n == 0 {
		return string(baseTo[0])
	}

	// Convert decimal to baseTo
	result := ""
	baseLen := len(baseTo)
	for n > 0 {
		result = string(baseTo[n%baseLen]) + result
		n /= baseLen
	}

	return result
}
