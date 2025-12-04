package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}

	var comb func([]int, int, int)
	comb = func(arr []int, start int, index int) {
		if index == n {
			for i := 0; i < n; i++ {
				z01.PrintRune(rune(arr[i] + '0'))
			}
			if arr[0] != 10-n {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			return
		}
		for i := start; i <= 9; i++ {
			arr[index] = i
			comb(arr, i+1, index+1)
		}
	}

	arr := make([]int, n)
	comb(arr, 0, 0)
	z01.PrintRune('\n')
}
