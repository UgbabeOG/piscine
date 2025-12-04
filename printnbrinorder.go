package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}

	var count [10]int
	if n == 0 {
		count[0] = 1
	}
	num := n
	for num > 0 {
		digit := num % 10
		count[digit]++
		num /= 10
	}

	for i := 0; i <= 9; i++ {
		for j := 0; j < count[i]; j++ {
			z01.PrintRune(rune('0' + i))
		}
	}
}

