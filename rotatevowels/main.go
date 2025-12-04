package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	// If no arguments → print newline only
	if len(args) == 0 {
		z01.PrintRune('\n')
		return
	}

	// Collect all vowels with their global positions
	type vowelInfo struct {
		pos  int
		char rune
	}
	var vowels []vowelInfo
	globalPos := 0

	for _, arg := range args {
		for _, r := range arg {
			if isVowel(r) {
				vowels = append(vowels, vowelInfo{globalPos, r})
			}
			globalPos++
		}
	}

	// If no vowels → print original arguments
	if len(vowels) == 0 {
		for i, arg := range args {
			if i > 0 {
				z01.PrintRune(' ')
			}
			for _, r := range arg {
				z01.PrintRune(r)
			}
		}
		z01.PrintRune('\n')
		return
	}

	// Build a map: position → mirrored vowel
	mirror := make(map[int]rune)
	for i := 0; i < len(vowels); i++ {
		left := vowels[i]
		right := vowels[len(vowels)-1-i]
		mirror[left.pos] = right.char
	}

	// Print all characters, replacing vowels with mirrored ones
	globalPos = 0
	argIndex := 0
	charIndex := 0

	for argIndex < len(args) {
		if argIndex > 0 {
			z01.PrintRune(' ')
		}
		arg := args[argIndex]
		charIndex = 0
		for charIndex < len([]rune(arg)) {
			currentChar := rune(arg[charIndex])
			if replacement, ok := mirror[globalPos]; ok {
				z01.PrintRune(replacement)
			} else {
				z01.PrintRune(currentChar)
			}
			globalPos++
			charIndex++
		}
		argIndex++
	}
	z01.PrintRune('\n')
}

func isVowel(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}
