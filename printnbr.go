package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	if n < 0 {
		z01.PrintRune('-')
		if n == -9223372036854775808 {
			s := "9223372036854775808"
			for _, r := range s {
				z01.PrintRune(r)
			}
			return
		}
		n = -n
	}

	var digits []int
	for n > 0 {
		digits = append(digits, n%10)
		n = n / 10
	}

	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(rune(digits[i] + '0'))
	}
}
