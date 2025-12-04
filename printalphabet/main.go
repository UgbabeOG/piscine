package main

import "github.com/01-edu/z01"

func main() {
	for k := 97; k < 123; k++ {
		z01.PrintRune(rune(k))
	}
	z01.PrintRune('\n')
}
