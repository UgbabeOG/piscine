package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 || containsHelpFlag(args) {
		printHelp()
		return
	}

	insertStr := ""
	order := false
	strArg := ""

	i := 0
	for i < len(args) {
		arg := args[i]
		if arg == "--order" || arg == "-o" {
			order = true
		} else if len(arg) >= 9 && arg[:9] == "--insert=" {
			insertStr = arg[9:]
		} else if arg == "-i" {
			if i+1 < len(args) {
				insertStr = args[i+1]
				i++
			}
		} else {
			strArg = arg
		}
		i++
	}

	result := strArg + insertStr

	if order {
		result = sortString(result)
	}

	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func containsHelpFlag(args []string) bool {
	for _, arg := range args {
		if arg == "--help" || arg == "-h" {
			return true
		}
	}
	return false
}

func printHelp() {
	printlnFlag("--insert", "  -i", "This flag inserts the string into the string passed as argument.")
	printlnFlag("--order", "  -o", "This flag will behave like a boolean, if it is called it will order the argument.")
}

func printlnFlag(long, short, desc string) {
	for _, r := range long {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
	for _, r := range short {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
	for i := 0; i < 9; i++ {
		z01.PrintRune(' ')
	}
	for _, r := range desc {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func sortString(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		for j := i + 1; j < len(runes); j++ {
			if runes[i] > runes[j] {
				runes[i], runes[j] = runes[j], runes[i]
			}
		}
	}
	return string(runes)
}
