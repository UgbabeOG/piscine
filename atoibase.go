package piscine

func AtoiBase(s string, base string) int {
	if len(base) < 2 {
		return 0
	}

	seen := make(map[rune]bool)
	for _, c := range base {
		if c == '+' || c == '-' {
			return 0
		}
		if seen[c] {
			return 0
		}
		seen[c] = true
	}

	digitValue := make(map[rune]int)
	for i, c := range base {
		digitValue[c] = i
	}

	sign := 1
	i := 0
	if len(s) > 0 {
		if s[0] == '-' {
			sign = -1
			i = 1
		} else if s[0] == '+' {
			i = 1
		}
	}

	result := 0
	n := len(base)
	for ; i < len(s); i++ {
		c := rune(s[i])
		value, ok := digitValue[c]
		if !ok {
			return 0
		}
		result = result*n + value
	}

	return result * sign
}
