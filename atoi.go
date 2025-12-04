package piscine

func Atoi(s string) int {
	result := 0
	sign := 1
	i := 0
	n := len(s)

	if n == 0 {
		return 0
	}

	if s[0] == '+' {
		i = 1
	} else if s[0] == '-' {
		sign = -1
		i = 1
	}

	for i < n {
		c := s[i]
		if c < '0' || c > '9' {
			return 0
		}
		digit := int(c - '0')
		result = result*10 + digit
		i++
	}
	return result * sign
}
