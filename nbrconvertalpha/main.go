package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}

	upper := false
	if args[0] == "--upper" {
		upper = true
		args = args[1:]
	}

	for _, arg := range args {
		n := Atoi(arg)
		if n < 1 || n > 26 {
			z01.PrintRune(' ')
			continue
		}

		var r rune
		if upper {
			r = rune('A' + n - 1)
		} else {
			r = rune('a' + n - 1)
		}

		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func Atoi(s string) int {
	num := 0
	for _, r := range s {
		if r < '0' || r > '9' {
			return 0
		}
		num = num*10 + int(r-'0')
	}
	return num
}
