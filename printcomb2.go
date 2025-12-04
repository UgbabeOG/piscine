package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	numbers := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

	for i := 0; i <= 98; i++ {
		for j := i + 1; j <= 99; j++ {

			// Print first number
			z01.PrintRune(numbers[i/10])
			z01.PrintRune(numbers[i%10])

			z01.PrintRune(' ')

			// Print second number
			z01.PrintRune(numbers[j/10])
			z01.PrintRune(numbers[j%10])

			// Separate numbers
			if i != 98 || j != 99 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
